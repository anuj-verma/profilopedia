package main

import (
	"fmt"

	config "github.com/anuj-verma/profilopedia/config"
	services "github.com/anuj-verma/profilopedia/services"
)


func main() {
	fmt.Println("Loading app configuration file...")
	config.LoadConfig("./config")

	facebookClient := services.GetFacebookClient(config.FacebookAppId(), config.FacebookAppSecret())
	users := facebookClient.SearchByQuery("varsha yadav")

	for _, user := range users {
		userProfile := facebookClient.GetUserInfo(user.Id)
		fmt.Println(userProfile)
	}
}
