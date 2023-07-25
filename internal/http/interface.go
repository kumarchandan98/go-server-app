package http

import "net/http"

type EmployeeHandler interface {
	GetById(http.ResponseWriter, *http.Request)

	GetAll(http.ResponseWriter, *http.Request)

	Create(http.ResponseWriter, *http.Request)

	Update(http.ResponseWriter, *http.Request)

	Delete(http.ResponseWriter, *http.Request)
}

type QuotesHandler interface {
	GetById(http.ResponseWriter, *http.Request)

	GetAll(http.ResponseWriter, *http.Request)

	Create(http.ResponseWriter, *http.Request)

	Update(http.ResponseWriter, *http.Request)

	Delete(http.ResponseWriter, *http.Request)
}
