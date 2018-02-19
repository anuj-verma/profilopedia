package profilopedia

import (
	"github.com/spf13/viper"
)

var appConfig *Config

type Config struct {
	port           string
	github_api_url string
}

func LoadConfig(configPaths ...string) {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	viper.ReadInConfig()
	appConfig = &Config{
		port:           viper.GetString("PORT"),
		github_api_url: viper.GetString("GITHUB_API_URL"),
	}
}

func Port() string {
	return appConfig.port
}

func GithubApiUrl() string {
	return appConfig.github_api_url
}
