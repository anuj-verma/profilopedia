package profilopedia

import(
  "fmt"

  fb "github.com/huandu/facebook"
)

var accessToken string

type UserProfile struct {
  ExternalId string `json:"id"`
  Name string `json:"name"`
  Emails []string `json:"emails"`
  Phone string `json:"phone"`
  Birthday string `json:"birthday"`
  Location UserLocation `json:"location"`
  Cover ProfileCover `json:"cover"`
}

type FacebookClient struct {
  accessToken string
}

type ProfileCover struct {
  Id string `json:"id"`
  CoverId string `json:"cover_id"`
  OffsetX string `json:"offset_x"`
  OffsetY string `json:"offset_y"`
  ImageUrl string `json:"source"`
}

type UserLocation struct {
  City string `json:"city"`
  Country string `json:country"`
  Zip string `json:"zip"`
}

type SearchResult struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

func GetFacebookClient(appId, appSecret string) FacebookClient {
  app := fb.New(appId, appSecret)

  facebookClient := FacebookClient{
    accessToken: app.AppAccessToken(),
  }
  return facebookClient
}

func (facebookClient *FacebookClient) GetUserInfo(userId string) UserProfile {
  url := "/" + userId
  res, _ := fb.Get(url, fb.Params{
    "fields": "id,name,birthday,phone,emails,location,cover",
    "access_token": facebookClient.accessToken,
  })
  var userProfile UserProfile
  err := res.DecodeField("", &userProfile)
  if err != nil {
      fmt.Printf("An error has happened 1 %v", err)
  }
  return userProfile
}

func (facebookClient *FacebookClient) SearchByQuery(query string) []SearchResult {
  res, _ := fb.Get("/search", fb.Params{
        "access_token": facebookClient.accessToken,
        "type":         "page",
        "q":            query,
    })

  var items []SearchResult
  err := res.DecodeField("data", &items)

  if err != nil {
      fmt.Printf("An error has happened %v", err)
  }
  return items
}
