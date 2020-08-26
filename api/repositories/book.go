package repositories

import (
	"api-example/api/models"
	"api-example/db"
)

type BookFilter struct {
	Title string
}

func GetAllBooks(filter BookFilter) *[]models.Book {
	var books []models.Book

	connection := db.Connection

	if filter.Title != "" {
		connection = connection.Where("title LIKE ?", "%" + filter.Title + "%")
	}

	connection.Find(&books)
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

func DeleteBook(ID int) *models.Book {
	var book models.Book
	err := db.Connection.First(&book, ID).Error
	if err != nil {
		return nil
	}
	db.Connection.Delete(&book)
	return &book
}