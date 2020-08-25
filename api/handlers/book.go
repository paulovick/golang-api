package handlersbook

import (
	"api-example/api/models"
	"api-example/api/repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	books := repositories.GetAllBooks()
	responseContent, err := json.Marshal(books)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseContent)
	}
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

	responseContent, err := json.Marshal(newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error serializing the book"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseContent)
}

func RegisterBook(router *mux.Router) {
	router.HandleFunc("/books", getAll).Methods(http.MethodGet)
	router.HandleFunc("/books", create).Methods(http.MethodPost)
}