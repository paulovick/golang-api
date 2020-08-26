package router

import (
	handlersbook "api-example/api/handlers"
	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	handlersbook.RegisterBook(api)
	handlersbook.RegisterAuthor(api)
}