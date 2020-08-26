package handlersbook

import (
	"api-example/api/models"
	"api-example/api/repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	books := repositories.GetAllBooks()
	responseContent, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseContent)
}

func getById(w http.ResponseWriter, r *http.Request) {
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

	responseContent, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseContent)
}

func create(w http.ResponseWriter, r *http.Request) {
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

func RegisterBook(router *mux.Router) {
	router.HandleFunc("/books", getAll).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookID}", getById).Methods(http.MethodGet)
	router.HandleFunc("/books", create).Methods(http.MethodPost)
}