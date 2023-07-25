package employee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-server-app/internal/models"
	"go-server-app/internal/services"

	"github.com/gorilla/mux"
)

type Employee struct {
	empService services.EmployeeInterface
}

type HTTPResponse struct {
	Message  string
	Success  bool
	Employee *models.Employee `json:"Employee,omitempty"`
}

func New(s services.EmployeeInterface) Employee {
	return Employee{empService: s}
}
func (e Employee) GetById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	emp, err := e.empService.GetEmployeeById(id)
	var res []byte
	if err != nil {
		response = HTTPResponse{Message: "Employee not found  : " + err.Error(), Success: false}
		res, _ = json.Marshal(response)
	} else {
		res, _ = json.Marshal(emp)
	}
	w.WriteHeader(200)
	w.Write([]byte(res))
}

func (e Employee) GetAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	var res []byte
	emp, err := e.empService.GetAllEmployees()
	if err != nil {
		response = HTTPResponse{Message: "Employee not found", Success: false}
		res, _ = json.Marshal(response)
	} else {
		res, _ = json.Marshal(emp)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (e Employee) Create(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response HTTPResponse

	decoder := json.NewDecoder(req.Body)
	var emp models.Employee
	err := decoder.Decode(&emp)
	if err != nil {
		w.Write([]byte("Error Parsing Body"))
		return
	}

	newemp, err := e.empService.CreateEmployee(&emp)
	fmt.Println("EmpCreate Called ", newemp)
	if err != nil {
		// w.Write([]byte(fmt.Sprintf("Error Creating Employee : %v", err)))
		response = HTTPResponse{Message: "Error Creating Employee : " + err.Error(), Success: false}
	} else {
		response = HTTPResponse{Message: "Employee Created", Success: true, Employee: newemp}
	}

	res, _ := json.Marshal(response)
	//fmt.Println(response, res)
	w.Write(res)
}

func (e Employee) Update(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	var emp models.Employee
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&emp)
	fmt.Println("Updated employee ", emp)
	if err != nil {
		// fmt.Println(err)
		w.Write([]byte("Error Parsing Body"))
		return
	}

	emp.Id = id

	updatedEmp, err := e.empService.UpdateEmployee(&emp)

	if err != nil {
		response = HTTPResponse{Message: "Cannot Update Employee : " + err.Error(), Success: false}

	} else {
		response = HTTPResponse{Message: "Employee Updated Successfully", Success: true, Employee: updatedEmp}
	}
	res, _ := json.Marshal(response)
	w.Write(res)

}

func (e Employee) Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response HTTPResponse
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	err := e.empService.DeleteEmployee(id)
	if err != nil {
		response = HTTPResponse{Message: "Cannot Delete Employee : " + err.Error(), Success: false}

	} else {
		response = HTTPResponse{Message: "Employee Deleted Successfully", Success: true}
	}
	res, _ := json.Marshal(response)
	w.Write(res)

}
