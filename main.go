package main

import (
	"fmt"

	"github.com/jadson-medeiros/go-rest/models"
	"github.com/jadson-medeiros/go-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Name One", About: "About One"},
		{Name: "Name Two", About: "About two"},
	}

	fmt.Println("Server started")
	routes.HandleRequest()
}
