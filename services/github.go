package profilopedia

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type githubRequest struct {
	Query   string
	Page    int
	PerPage int
}

type githubResponse struct {
	TotalCount        int                    `json:"total_count"`
	IncompleteResults bool                   `json:"incomplete_results"`
	Items             []githubResponseFields `json:"items"`
}

type githubResponseFields struct {
	Login             string  `json:"login"`
	Id                int     `json:"id"`
	AvatarUrl         string  `json:"avatar_url"`
	GravatarId        string  `json:"gravatar_id"`
	Url               string  `json:"url"`
	HtmlUrl           string  `json:"html_url"`
	FollowersUrl      string  `json:"followers_url"`
	FollowingUrl      string  `json:"following_url"`
	GistsUrl          string  `json:"gists_url"`
	StarredUrl        string  `json:"starred_url"`
	SubscriptionsUrl  string  `json:"subscriptions_url"`
	OrganizationsUrl  string  `json:"organizations_url"`
	ReposUrl          string  `json:"repos_url"`
	EventsUrl         string  `json:"events_url"`
	ReceivedEventsUrl string  `json:"received_events_url"`
	SiteAdmin         bool    `json:"site_admin"`
	Type              string  `json:"type"`
	Score             float64 `json:"score"`
}

type githubClient struct {
	ApiUrl   string
	Request  *githubRequest
	Response *githubResponse
}

func (client *githubClient) Search(query string, page, perPage int) {
	client.PrepareGithubRequest(query, page, perPage)
	url := fmt.Sprintf("%s?q=%s&page=%d&per_page=%d", client.ApiUrl, client.Request.Query, client.Request.Page, client.Request.PerPage)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer res.Body.Close()
	var data githubResponse
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Println(err)
	}
	client.Response = &data
}

func GetGithubClient(apiUrl string) githubClient {
	return githubClient{ApiUrl: apiUrl}
}

func (client *githubClient) PrepareGithubRequest(query string, page int, perPage int) {
	client.Request = &githubRequest{
		Query:   query,
		Page:    page,
		PerPage: perPage,
	}
}
