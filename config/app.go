package profilopedia

import (
	"github.com/spf13/viper"
)

var appConfig *Config

type Config struct {
	Port                     string
	GithubApiUrl             string
	DefaultPage              int
	DefaultPerPage           int
	TwitterConsumerKey       string
	TwitterConsumerSecret    string
	TwitterAccessToken       string
	TwitterAccessTokenSecret string
	DbUser                   string
	DbName                   string
	DbPassword               string
}

func LoadConfig(configPaths ...string) {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	viper.ReadInConfig()
	appConfig = &Config{
		Port:                     viper.GetString("PORT"),
		GithubApiUrl:             viper.GetString("GITHUB_API_URL"),
		DefaultPage:              viper.GetInt("DEFAULT_PAGE"),
		DefaultPerPage:           viper.GetInt("DEFAULT_PER_PAGE"),
		TwitterConsumerKey:       viper.GetString("TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret:    viper.GetString("TWITTER_CONSUMER_SECRET"),
		TwitterAccessToken:       viper.GetString("TWITTER_ACCESS_TOKEN"),
		TwitterAccessTokenSecret: viper.GetString("TWITTER_ACCESS_TOKEN_SECRET"),
		DbUser:     viper.GetString("DB_USER"),
		DbName:     viper.GetString("DB_NAME"),
		DbPassword: viper.GetString("DB_PASSWORD"),
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

func TwitterConsumerKey() string {
	return appConfig.TwitterConsumerKey
}

func TwitterConsumerSecret() string {
	return appConfig.TwitterConsumerSecret
}

func TwitterAccessToken() string {
	return appConfig.TwitterAccessToken
}

func TwitterAccessTokenSecret() string {
	return appConfig.TwitterAccessTokenSecret
}

func DbName() string {
	return appConfig.DbName
}

func DbUser() string {
	return appConfig.DbUser
}

func DbPassword() string {
	return appConfig.DbPassword
}
