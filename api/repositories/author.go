package repositories

import (
	"api-example/api/models"
	"api-example/db"
)

type AuthorFilter struct {
	Name string
}

func GetAllAuthors(filter AuthorFilter) *[]models.Author {
	var authors []models.Author

	connection := db.Connection

	if filter.Name != "" {
		connection = connection.Where("name LIKE ?", "%" + filter.Name + "%")
	}

	connection.Find(&authors)
	return &authors
}

func GetAuthorById(ID int) *models.Author {
	var author models.Author
	err := db.Connection.First(&author, ID).Error
	if err != nil {
		return nil
	}
	return &author
}

func CreateAuthor(author *models.Author) *models.Author {
	db.Connection.Create(author)
	return author
}

func DeleteAuthor(ID int) *models.Author {
	var author models.Author
	err := db.Connection.First(&author, ID).Error
	if err != nil {
		return nil
	}
	db.Connection.Delete(&author)
	return &author
}