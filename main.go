package main

import (
	"fmt"
	"sync"

	appConfig "github.com/anuj-verma/profilopedia/config"
	services "github.com/anuj-verma/profilopedia/services"
	"github.com/dghubble/go-twitter/twitter"
)

func fetchGithubData(wg *sync.WaitGroup) {
	// Decrement the counter when the goroutine completes.
	defer wg.Done()
	githubClient := services.GetGithubClient(appConfig.GithubApiUrl())
	githubClient.Search("anuj", appConfig.DefaultPage(), appConfig.DefaultPerPage())
	data := githubClient.Response
	for _, value := range data.Items {
		fmt.Println("Github: ", value.Login)
	}

}

func fetchTwitterData(wg *sync.WaitGroup) {
	// Decrement the counter when the goroutine completes.
	defer wg.Done()
	twitterClient := services.GetTwitterClient(
		appConfig.TwitterConsumerKey(),
		appConfig.TwitterConsumerSecret(),
		appConfig.TwitterAccessToken(),
		appConfig.TwitterAccessTokenSecret(),
	)
	search, _, _ := twitterClient.Users.Search("anuj", &twitter.UserSearchParams{
		Query: "anuj",
	})
	for _, value := range search {
		fmt.Println("Twitter: ", value.Name)
	}
}
func main() {
	fmt.Println("Loading app configuration file...")
	appConfig.LoadConfig("./config")

	var wg sync.WaitGroup
	// Increment the WaitGroup counter to number of go routines
	wg.Add(2)
	// Call go routines
	go fetchGithubData(&wg)
	go fetchTwitterData(&wg)
	// Wait for all go routines to complete
	wg.Wait()
}
