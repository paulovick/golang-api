package database

import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	DBConnection *gorm.DB
)

func InitDB() {
	var err error
	DBConnection, err = gorm.Open("postgres", "host=localhost port=5432 user=apiexample dbname=apiexample password=apiexample sslmode=disable")

	if err != nil {
		log.Fatal("Error connection to the database", err)
	}
}

func CloseDB() {
	defer DBConnection.Close()
}