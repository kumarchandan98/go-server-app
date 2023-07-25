package employee

import (
	"errors"
	"fmt"
	"go-server-app/internal/metrics"
	"go-server-app/internal/models"
)

type EmployeeService struct {
	cache map[int]*models.Employee
}

func New() EmployeeService {

	return EmployeeService{cache: map[int]*models.Employee{}}
}

func (es *EmployeeService) CreateEmployee(emp *models.Employee) (*models.Employee, error) {
	if emp.Name == "" || emp.Email == "" || emp.Role == "" {
		return nil, errors.New("fields Cannot be Empty")
	}
	if !ValidateEmail(emp.Email) {
		return nil, errors.New("invalid email id")
	}
	if !ValidateRole(emp.Role) {
		return nil, errors.New("role must be one of [developer,manager]")
	}

	if _, ok := es.cache[emp.Id]; ok {
		return nil, errors.New("ID already exists")
	}

	es.cache[emp.Id] = emp
	metrics.IncreaseEmployee(emp.Id)
	fmt.Println("Incremented Metrics")
	return emp, nil
}

func (es *EmployeeService) GetEmployeeById(id int) (*models.Employee, error) {
	if id < 0 {
		return nil, errors.New("id cannot be negative")
	}

	emp, ok := es.cache[id]
	if !ok {
		return nil, errors.New("ID already exists")
	}

	return emp, nil
}

func (es *EmployeeService) GetAllEmployees() ([]*models.Employee, error) {
	employees := []*models.Employee{}
	for _, emp := range es.cache {
		employees = append(employees, emp)
	}
	metrics.GetEmployee(len(employees))
	return employees, nil
}

func (es *EmployeeService) UpdateEmployee(emp *models.Employee) (*models.Employee, error) {
	_, ok := es.cache[emp.Id]
	if !ok {
		return nil, errors.New("ID does not exists")
	}

	es.cache[emp.Id] = emp

	return emp, nil
}

func (es *EmployeeService) DeleteEmployee(id int) error {
	if id < 0 {
		return errors.New("id cannot be negative")
	}

	emp, ok := es.cache[id]
	if !ok {
		return errors.New("ID does not exists")
	}

	delete(es.cache, emp.Id)

	return nil
}
