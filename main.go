package main

import (
	"fmt"
	"github.com/anuj-verma/profilopedia/app"
)

func main() {
	fmt.Println("Loading app configuration file...")
	app.LoadConfig("./config")
	fmt.Println("Port from config is =>", app.Port())
}
