package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jadson-medeiros/go-rest/controllers"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalities", controllers.GetAllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", router))
}
