package main

import (
	"log"
	"net/http"

	"github.com/guisantosalves/API-Golang-mysql/models"
	"github.com/guisantosalves/API-Golang-mysql/router"
)

func main() {

	models.InitialMigration()
	log.Fatal(http.ListenAndServe(":9000", router.InitializeRouter()))

}
