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

func GetBookById(ID int) *models.Book {
	var book models.Book
	err := db.Connection.First(&book, ID).Error
	if err != nil {
		return nil
	}
	return &book
}

func CreateBook(book *models.Book) *models.Book {
	db.Connection.Create(book)
	return book
}