package routes

import (
	handler "go-server-app/internal/http/employee"

	"github.com/gorilla/mux"
)

func SetEmployeeRoutes(router *mux.Router, handler handler.Employee) {
	router.HandleFunc("", handler.GetAll).Methods("GET")

	router.HandleFunc("/{id}", handler.GetById).Methods("GET")

	router.HandleFunc("", handler.Create).Methods("POST")

	router.HandleFunc("/{id}", handler.Update).Methods("PUT")

	router.HandleFunc("/{id}", handler.Delete).Methods("DELETE")

}
