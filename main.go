package main

import (
	"fmt"

	"github.com/jadson-medeiros/go-rest/routes"
)

func main() {
	fmt.Println("Server started")
	routes.HandleRequest()
}
