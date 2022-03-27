package config

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// v allows access to the underlying viper instance if more advanced getters are needed
var v = setup()

// setup sets up viper
func setup() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	if err != nil {
		panic("Unable to load configuration: " + err.Error())
	}

	// Set environment variable overrides
	if v.IsSet("app.prefix") {
		// An environment prefix is set
		v.SetEnvPrefix(v.GetString("app.prefix"))
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		v.AutomaticEnv()
	}

	// Init the logger
	log, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(log)
	zap.L().Info("Started logger")

	return v
}

// GetString return value of string config param
func GetString(key string) string {
	return v.GetString(key)
}

// GetInt return value of int config param
func GetInt(key string) int {
	return v.GetInt(key)
}
