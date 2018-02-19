package profilopedia

import "fmt"

type githubRequest struct {
	query    string
	sort     string
	order    string
	page     int
	per_page int
}

type githubResponse struct {
	login      string
	id         int
	avatar_url string
	url        string
	score      string
}

type githubClient struct {
	api_url string
}

func (client githubClient) Search() (response []githubResponse) {
	fmt.Println("Searching the github api...")
	return nil
}

func GetGithubClient(api_url string) githubClient {
	return githubClient{api_url: api_url}
}
