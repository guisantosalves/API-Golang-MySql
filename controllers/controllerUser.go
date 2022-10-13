package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guisantosalves/API-Golang-mysql/models"
)

// w -> response // r -> request
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	// array of users
	user := []models.User{}

	// getting all the results
	models.DB.Find(&user)

	json.NewEncoder(w).Encode(user)
}

// w response, r request
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting params from url
	params := mux.Vars(r)

	var user models.User

	models.DB.First(&user, params["id"])

	json.NewEncoder(w).Encode(user)
}

// working
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// setting in header
	w.Header().Set("Content-Type", "application/json")

	// decoding in user structure type
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	//creating
	// get from r.body and put in DB.Create
	models.DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	// get params var
	params := mux.Vars(r)

	// find the row that I want to change
	models.DB.First(&user, params["id"])

	// get from body and set in &user
	json.NewDecoder(r.Body).Decode(&user)

	// then &user from db is overwrited to &user from body
	models.DB.Save(&user)

	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	idFromParams := mux.Vars(r)

	models.DB.Delete(&user, idFromParams["id"])

	json.NewEncoder(w).Encode("The user was deleted.")
}
