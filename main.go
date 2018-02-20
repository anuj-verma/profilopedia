package main

import (
	"fmt"

	config "github.com/anuj-verma/profilopedia/config"
	services "github.com/anuj-verma/profilopedia/services"
)

func main() {
	fmt.Println("Loading app configuration file...")
	config.LoadConfig("./config")

  accessToken := config.FacebookAccessToken()
  fbClient := services.GetFacebookClient(accessToken)
  items := fbClient.UserSearch("varsha yadav")
  for _, item := range items {
		userProfile := fbClient.GetUserInfo(item.Id)
		fmt.Println(userProfile)
  }
}
