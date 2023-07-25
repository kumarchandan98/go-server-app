package routes

import (
	handler "go-server-app/internal/http/quotes"

	"github.com/gorilla/mux"
)

func SetQuotesRoutes(router *mux.Router, handler handler.Quotes) {

	// get all quotes by all employees
	router.HandleFunc("", handler.GetAll).Methods("GET")

	// get all quotes by employee id
	router.HandleFunc("/{empId}", handler.GetById).Methods("GET")

	// Post quotes by employee id
	router.HandleFunc("", handler.Create).Methods("POST")

	// Update quote by employee id
	router.HandleFunc("/{empId}/{quoteId}", handler.Update).Methods("PUT")

	// Delete quote by employee id
	router.HandleFunc("/{empId}", handler.Delete).Methods("DELETE")

}
