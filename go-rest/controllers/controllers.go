package controllers

import (
	"encoding/json"
	"fmt"
	"goRest/database"
	"goRest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidade []models.Personalidade

	database.DB.Find(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func GetPersonalidadeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func CreatePersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Save(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}
