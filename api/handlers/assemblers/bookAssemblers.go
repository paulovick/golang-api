package assemblers

import (
	apimodels "api-example/api/handlers/models"
	"api-example/api/models"
	"api-example/api/repositories"
)

func AssembleBookResponse(book *models.Book) *apimodels.BookResponse {
	author := repositories.GetAuthorById(book.AuthorID)

	return &apimodels.BookResponse{
		ID:            book.ID,
		Title:         book.Title,
		YearOfPublish: book.YearOfPublish,
		Author:		   apimodels.BookAuthorResponse{
			ID:   author.ID,
			Name: author.Name,
		},
	}
}
