package db

import (
	"api-example/api/models"
	"api-example/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var logger = customlogger.GetLogger("MIGRATIONS")

func autoMigrate() {
	Connection.AutoMigrate(&models.Book{})
}

func applyMigrations() {
	driver, err := postgres.WithInstance(Connection.DB(), &postgres.Config{})
	if err != nil {
		logger.Fatal("Error connecting to postgres for applying migrations:", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "apiexample", driver)
	if err != nil {
		logger.Fatal("Error obtaining migrations:", err)
	}

	logger.Info("Applying migrations...")

	err = m.Up()
	if err != nil {
		logger.Fatal("Error applying migrations:", err)
	}

	logger.Info("Migrations applied successfully")
}

func Migrate() {
	// autoMigrate()
	applyMigrations()
}