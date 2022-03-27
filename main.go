package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"lead_generation_basic/config"
)

func main() {
	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetOutput(ioutil.Discard)

	// Routes
	Routes(e)

	// Start the server
	go func() {
		zap.L().Info("Starting HTTP server", zap.String("address", config.GetString("http.address")))
		err := e.Start(config.GetString("http.address"))
		if err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("Shutting down the server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	err := e.Shutdown(ctx)
	zap.L().Info("Server shutdown")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
