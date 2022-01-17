package routes

import (
	"log"
	"net/http"

	"github.com/jadson-medeiros/go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
