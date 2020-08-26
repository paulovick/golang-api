package handlers

import (
	"api-example/api/handlers/assemblers"
	apimodels "api-example/api/handlers/models"
	"api-example/api/models"
	"api-example/api/repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := repositories.BookFilter{
		Title: query.Get("title"),
	}

	books := *repositories.GetAllBooks(filter)

	var bookResponses []apimodels.BookResponse
	for i := 0; i < len(books); i++ {
		book := books[i]
		author := repositories.GetAuthorById(book.AuthorID)
		bookResponses = append(bookResponses, *assemblers.AssembleBookResponse(&book, author))
	}

	responseContent, _ := json.Marshal(bookResponses)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseContent)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := strconv.Atoi(params["bookID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"bookID" should be a number`))
		return
	}

	book := repositories.GetBookById(bookID)

	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	author := repositories.GetAuthorById(book.AuthorID)
	bookResponse := assemblers.AssembleBookResponse(book, author)
	responseContent, _ := json.Marshal(bookResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseContent)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decoding the request"))
		return
	}

	newBook := repositories.CreateBook(&book)

	responseContent, _ := json.Marshal(newBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseContent)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := strconv.Atoi(params["bookID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"bookID" should be a number`))
		return
	}

	removedBook := repositories.DeleteBook(bookID)

	if removedBook == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseContent, _ := json.Marshal(removedBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseContent)
}

func RegisterBook(router *mux.Router) {
	router.HandleFunc("/books", getAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookID}", getBookById).Methods(http.MethodGet)
	router.HandleFunc("/books", createBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{bookID}", deleteBook).Methods(http.MethodDelete)
}