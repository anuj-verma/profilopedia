package main

import (
	"fmt"

	config "github.com/anuj-verma/profilopedia/config"
	services "github.com/anuj-verma/profilopedia/services"
)

func main() {
	fmt.Println("Loading app configuration file...")
	config.LoadConfig("./config")

	fmt.Println(config.GithubApiUrl())
	github_client := services.GetGithubClient(config.GithubApiUrl())
	github_client.Search("anuj")
}
