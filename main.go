package main

import (
	"context"
	"fmt"
	empHandler "go-server-app/internal/http/employee"
	qtHandler "go-server-app/internal/http/quotes"
	"go-server-app/internal/metrics"
	"go-server-app/internal/routes"
	empSrv "go-server-app/internal/services/employee"
	qtSrv "go-server-app/internal/services/quotes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mp, err := metrics.GetMeter()
	if err != nil {
		fmt.Errorf("error when starting meter %s", err)
	}
	defer mp.Shutdown(context.Background())

	empService := empSrv.New()
	employeeHandler := empHandler.New(&empService)

	quoteService := qtSrv.New()
	quoteHandler := qtHandler.New(quoteService)

	router := mux.NewRouter()
	empPrefix := router.PathPrefix("/employee").Subrouter()
	quotesPrefix := router.PathPrefix("/quotes").Subrouter()

	routes.SetEmployeeRoutes(empPrefix, employeeHandler)
	routes.SetQuotesRoutes(quotesPrefix, quoteHandler)

	http.Handle("/", router)
	fmt.Println("Server started http://localhost:28080")

	if err := http.ListenAndServe(":28080", nil); err != nil {
		fmt.Println("Error ", err)
	}

}
