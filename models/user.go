package models

import (
	"os"

	"gorm.io/gorm"
)

var stringConnection = os.Getenv("CONNECTIONSTR")

var DB *gorm.DB
var err error

var DNS = stringConnection

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {

}
