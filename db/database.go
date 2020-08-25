package db

import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	Connection *gorm.DB
)

func InitDB() {
	var err error
	Connection, err = gorm.Open("postgres", "host=localhost port=5432 user=apiexample dbname=apiexample password=apiexample sslmode=disable")

	if err != nil {
		log.Fatal("Error connection to the db", err)
	}
}

func CloseDB() {
	defer Connection.Close()
}