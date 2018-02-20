package profilopedia

import (
	"github.com/spf13/viper"
)

var appConfig *Config

type Config struct {
	Port           string
	GithubApiUrl   string
	DefaultPage    int
	DefaultPerPage int
	FacebookAccessToken string
}

func LoadConfig(configPaths ...string) {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	viper.ReadInConfig()
	appConfig = &Config{
		Port:           viper.GetString("PORT"),
		GithubApiUrl:   viper.GetString("GITHUB_API_URL"),
		DefaultPage:    viper.GetInt("DEFAULT_PAGE"),
		DefaultPerPage: viper.GetInt("DEFAULT_PER_PAGE"),
		FacebookAccessToken: viper.GetString("FACEBOOK_ACCESS_TOKEN"),
	}
}

func Port() string {
	return appConfig.Port
}

func GithubApiUrl() string {
	return appConfig.GithubApiUrl
}

func DefaultPage() int {
	return appConfig.DefaultPage
}

func DefaultPerPage() int {
	return appConfig.DefaultPerPage
}

func FacebookAccessToken() string {
	return appConfig.FacebookAccessToken
}
