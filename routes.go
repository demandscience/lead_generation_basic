package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"lead_generation_basic/data"
)

// Routes bind endpoints
func Routes(e *echo.Echo) {
	e.GET("/sample", sample)
}

// sample get all endpoint
func sample(c echo.Context) error {
	res, _ := data.GetAll()
	return c.JSON(http.StatusOK, res)
}
