package main

import (
	router "api-example/api"
	"api-example/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db.InitDB()
	db.Migrate()

	r := mux.NewRouter()
	router.SetRoutes(r)

	err := http.ListenAndServe(":8080", r)

	db.CloseDB()

	if err != nil {
		log.Fatal(err)
	}
}