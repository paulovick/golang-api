package handlersbook

import (
	"github.com/gorilla/mux"
	"net/http"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get all books"}`))
}

func RegisterBook(router *mux.Router) {
	router.HandleFunc("/books", getAll)
}