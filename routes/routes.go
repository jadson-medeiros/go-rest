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
	router.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("Get")
	router.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("Get")
	router.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", router))
}
