package repositories

import (
	"api-example/api/models"
	"api-example/db"
)

func GetAllBooks() *[]models.Book {
	var books []models.Book
	db.Connection.Find(&books)
	return &books
}

func CreateBook(book *models.Book) *models.Book {
	db.Connection.Create(book)
	return book
}