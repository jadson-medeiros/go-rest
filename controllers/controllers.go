package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jadson-medeiros/go-rest/database"
	"github.com/jadson-medeiros/go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality

	database.DB.Table("personality").Find(&p)

	json.NewEncoder(w).Encode(p)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.Table("personality").First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)

	database.DB.Table("personality").Create(&newPersonality)

	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.Table("personality").Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.Table("personality").First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Table("personality").Save(&personality)

	json.NewEncoder(w).Encode(personality)
}
