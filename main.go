package main

import (
	"fmt"

	config "github.com/anuj-verma/profilopedia/config"
	services "github.com/anuj-verma/profilopedia/services"
)

func main() {
	fmt.Println("Loading app configuration file...")
	config.LoadConfig("./config")

	githubClient := services.GetGithubClient(config.GithubApiUrl())
	githubClient.Search("anuj", config.DefaultPage(), config.DefaultPerPage())
	data := githubClient.Response
	for _, value := range data.Items {
		fmt.Println(value)
	}
}
