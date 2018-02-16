package app

import (
	"github.com/spf13/viper"
)

var appConfig *Config

type Config struct {
	port string
}

func LoadConfig(configPaths ...string) {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	viper.ReadInConfig()
	appConfig = &Config{
		port: viper.GetString("PORT"),
	}
}

func Port() string {
	return appConfig.port
}
