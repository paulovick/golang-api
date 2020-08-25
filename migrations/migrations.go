package migrations

import (
	"api-example/api/models"
	"api-example/database"
)

func Migrate() {
	database.DBConnection.AutoMigrate(&models.Book{})
}