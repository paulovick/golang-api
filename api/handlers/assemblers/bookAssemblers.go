package assemblers

import (
	apimodels "api-example/api/handlers/models"
	"api-example/api/models"
)

func AssembleBookResponse(book *models.Book, author *models.Author) *apimodels.BookResponse {
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
