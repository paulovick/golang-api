package handlers

import (
	"api-example/api/models"
	"api-example/api/repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := repositories.AuthorFilter{
		Name: query.Get("name"),
	}

	authors := repositories.GetAllAuthors(filter)
	responseContent, _ := json.Marshal(authors)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseContent)
}

func getAuthorById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	authorID, err := strconv.Atoi(params["authorID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"authorID" should be a number`))
		return
	}

	author := repositories.GetAuthorById(authorID)

	if author == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseContent, _ := json.Marshal(author)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseContent)
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decoding the request"))
		return
	}

	newAuthor := repositories.CreateAuthor(&author)

	responseContent, _ := json.Marshal(newAuthor)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseContent)
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	authorID, err := strconv.Atoi(params["authorID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"authorID" should be a number`))
		return
	}

	removedAuthor := repositories.DeleteAuthor(authorID)

	if removedAuthor == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseContent, _ := json.Marshal(removedAuthor)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseContent)
}

func RegisterAuthor(router *mux.Router) {
	router.HandleFunc("/authors", getAllAuthors).Methods(http.MethodGet)
	router.HandleFunc("/authors/{authorID}", getAuthorById).Methods(http.MethodGet)
	router.HandleFunc("/authors", createAuthor).Methods(http.MethodPost)
	router.HandleFunc("/authors/{authorID}", deleteAuthor).Methods(http.MethodDelete)
}