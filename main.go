package main

import (
	"api-example/api/handlers"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func getCommentFromUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	if val, ok := pathParams["userID"]; ok {
		var err error
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "userID must be a number"}`))
		}
	}

	commentID := -1
	if val, ok := pathParams["commendID"]; ok {
		var err error
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "commentID must be a number"}`))
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s"}`, userID, commentID, location)))
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	handlers_book.RegisterBook(api)

	log.Fatal(http.ListenAndServe(":8080", r))
}