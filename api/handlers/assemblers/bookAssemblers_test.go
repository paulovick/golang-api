package assemblers

import (
	apimodels "api-example/api/handlers/models"
	"api-example/api/models"
	testHelpers "api-example/test"
	"testing"
)

func TestAssembleBookResponse(t *testing.T) {
	book := models.Book{
		ID:            1,
		Title:         "Some Title",
		YearOfPublish: 1999,
		AuthorID:      2,
	}
	author := models.Author{
		ID:   2,
		Name: "Some name",
	}
	expected := apimodels.BookResponse{
		ID:            1,
		Title:         "Some Title",
		YearOfPublish: 1999,
		Author:        apimodels.BookAuthorResponse{
			ID:   2,
			Name: "Some name",
		},
	}

	actual := *AssembleBookResponse(&book, &author)

	testHelpers.AssertEqual(t, actual, expected, "Book response is not the correct one")
}
