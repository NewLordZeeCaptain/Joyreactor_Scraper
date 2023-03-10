package main

import (
	"Projects/config"
	"Projects/database/postgres"
	"Projects/services"
	"fmt"
)

func main() {
	config.LoadConfig()
	postgres.InitDB()

	content := make(map[string]string)
	services.Scraper(content)
	var UserName string
	fmt.Scanln(&UserName)
	user, err := services.GetUserInfo(UserName)
	if err != nil {
		panic(err)
	}
	postgres.InsertUser(user)

}
