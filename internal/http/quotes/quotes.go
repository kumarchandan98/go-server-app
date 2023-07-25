package quotes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-server-app/internal/models"
	"go-server-app/internal/services"

	"github.com/gorilla/mux"
)

type Quotes struct {
	quoteService services.QuotesInterface
}

type HTTPResponse struct {
	Message string
	Success bool
	Quotes  *models.Quote `json:"Quotes,omitempty"`
}

func New(s services.QuotesInterface) Quotes {
	return Quotes{quoteService: s}
}

func (e Quotes) GetById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["empId"])
	quoteId, _ := strconv.Atoi(vars["quoteId"])

	emp, err := e.quoteService.GetQuoteById(id, quoteId)
	var res []byte
	if err != nil {
		response = HTTPResponse{Message: "Quotes not found: " + err.Error(), Success: false}
		res, _ = json.Marshal(response)
	} else {
		res, _ = json.Marshal(emp)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (e Quotes) GetAllQuoteByEmpId(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	vars := mux.Vars(req)
	empId, _ := strconv.Atoi(vars["empId"])

	quotes, err := e.quoteService.GetAllQuoteByEmpId(empId)
	fmt.Println(quotes)
	if err != nil {
		response = HTTPResponse{Message: "Cannot get quotes for employee: " + err.Error(), Success: false}
	} else {
		response = HTTPResponse{Message: "Quotes for employee retrieved successfully", Success: true, Quotes: nil}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}

func (e Quotes) GetAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	var res []byte
	emp, err := e.quoteService.GetAllQuotes()
	if err != nil {
		response = HTTPResponse{Message: "Quotes not found", Success: false}
		res, _ = json.Marshal(response)
	} else {
		res, _ = json.Marshal(emp)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (e Quotes) Create(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response HTTPResponse

	decoder := json.NewDecoder(req.Body)
	var emp models.Quote
	err := decoder.Decode(&emp)
	if err != nil {
		w.Write([]byte("Error Parsing Body"))
		return
	}

	newemp, err := e.quoteService.CreateQuote(emp.Id, &emp) // Pass emp.Id as an argument
	fmt.Println("EmpCreate Called ", newemp)
	if err != nil {
		// w.Write([]byte(fmt.Sprintf("Error Creating Quotes : %v", err)))
		response = HTTPResponse{Message: "Error Creating Quotes: " + err.Error(), Success: false}
	} else {
		response = HTTPResponse{Message: "Quotes Created", Success: true, Quotes: newemp}
	}

	res, _ := json.Marshal(response)
	fmt.Println(response, res)
	w.Write(res)
}

func (e Quotes) Update(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	vars := mux.Vars(req)
	empId, _ := strconv.Atoi(vars["empId"])
	quoteId, _ := strconv.Atoi(vars["quoteId"])

	var updates map[string]string
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&updates)
	fmt.Println("Updates ", updates)
	if err != nil {
		// fmt.Println(err)
		w.Write([]byte("Error Parsing Body"))
		return
	}

	updatedEmp, err := e.quoteService.UpdateQuote(empId, quoteId, updates) // Pass empId and quoteId as arguments

	if err != nil {
		response = HTTPResponse{Message: "Cannot Update Quotes: " + err.Error(), Success: false}

	} else {
		response = HTTPResponse{Message: "Quotes Updated Successfully", Success: true, Quotes: updatedEmp}
	}
	res, _ := json.Marshal(response)
	w.Write(res)
}

func (e Quotes) Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	vars := mux.Vars(req)
	empId, _ := strconv.Atoi(vars["empId"])
	quoteId, _ := strconv.Atoi(vars["quoteId"])

	err := e.quoteService.DeleteQuote(empId, quoteId)
	if err != nil {
		response = HTTPResponse{Message: "Cannot Delete Quote: " + err.Error(), Success: false}
	} else {
		response = HTTPResponse{Message: "Quote Deleted Successfully", Success: true}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
