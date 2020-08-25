package db

import (
	"api-example/api/models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func autoMigrate() {
	Connection.AutoMigrate(&models.Book{})
}

func applyMigrations() {
	driver, err := postgres.WithInstance(Connection.DB(), &postgres.Config{})
	if err != nil {
		log.Fatal("Error connecting to postgres for applying migrations: ", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "apiexample", driver)
	if err != nil {
		log.Fatal("Error obtaining migrations: ", err)
	}

	log.Println("Applying migrations...")

	err = m.Up()
	if err != nil {
		log.Fatal("Error applying migrations: ", err)
	}

	log.Println("Migrations applied successfully")

	defer m.Close()
}

func Migrate() {
	// autoMigrate()
	applyMigrations()
}