package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest/database"
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
func AllPersonas(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}
func RetornPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}
func CreatePersona(w http.ResponseWriter, r *http.Request) {
	var newPersona models.Personalidade

	json.NewDecoder(r.Body).Decode(&newPersona)
	database.DB.Create(&newPersona)
	json.NewEncoder(w).Encode(newPersona)
}
func DeletePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}
func EditPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)

	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
