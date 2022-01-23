package main

import (
	"fmt"

	"github.com/jadson-medeiros/go-rest/database"
	"github.com/jadson-medeiros/go-rest/models"
	"github.com/jadson-medeiros/go-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name One", About: "About One"},
		{Id: 2, Name: "Name Two", About: "About two"},
	}

	database.ConnectionDB()

	fmt.Println("Server started")
	routes.HandleRequest()
}
