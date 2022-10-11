package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/guisantosalves/API-Golang-mysql/models"
)

// w -> response // r -> request
func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

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

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
