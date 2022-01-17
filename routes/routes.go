package routes

import (
	"log"
	"net/http"

	"github.com/jadson-medeiros/go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.GetAllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
