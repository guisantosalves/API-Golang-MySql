package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var stringConnection = os.Getenv("CONNECTIONSTR")

var DB *gorm.DB
var err error

var DNS = "root:15022002@/users"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to db")
	}

	DB.AutoMigrate(&User{})
}
