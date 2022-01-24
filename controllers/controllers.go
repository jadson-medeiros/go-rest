package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jadson-medeiros/go-rest/database"
	"github.com/jadson-medeiros/go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]

	// for _, personality := range models.Personalities {
	// 	if strconv.Itoa(personality.Id) == id {
	// 		json.NewEncoder(w).Encode(personality)
	// 	}
	// }
}
