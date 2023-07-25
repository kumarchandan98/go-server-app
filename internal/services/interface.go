package services

import "go-server-app/internal/models"

type EmployeeInterface interface {
	CreateEmployee(*models.Employee) (*models.Employee, error)
	GetEmployeeById(id int) (*models.Employee, error)
	GetAllEmployees() ([]*models.Employee, error)
	UpdateEmployee(*models.Employee) (*models.Employee, error)
	DeleteEmployee(id int) error
}

type QuotesInterface interface {
	CreateQuote(empId int, quote *models.Quote) (*models.Quote, error)
	GetQuoteById(int, int) (*models.Quote, error)
	GetAllQuotes() ([]models.Quote, error)
	GetAllQuoteByEmpId(empId int) ([]models.Quote, error)
	UpdateQuote(empId, quoteId int, updates map[string]string) (*models.Quote, error)
	DeleteQuote(empId, quoteId int) error
}
