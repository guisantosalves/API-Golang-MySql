package router

import (
	"github.com/gorilla/mux"
	"github.com/guisantosalves/API-Golang-mysql/controllers"
)

func InitializeRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	return r
}
