package main

import (
	"fmt"

	"github.com/jadson-medeiros/go-rest/database"
	"github.com/jadson-medeiros/go-rest/routes"
)

func main() {
	database.ConnectionDB()

	fmt.Println("Server started")
	routes.HandleRequest()
}
