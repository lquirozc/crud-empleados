package services

import (
	"crud-empleados/models"
	repository "crud-empleados/repositories"
)

// EmployeeService defines the methods that the service layer will provide.
type EmployeeService interface {
	CreateEmployee(employee *models.Employee) (*models.Employee, error)
	GetEmployeeByID(employeeID string) (*models.Employee, error)
	GetAllEmployees() ([]models.Employee, error)
	UpdateEmployee(employee *models.Employee) (*models.Employee, error)
	DeleteEmployee(employeeID string) error
}

// employeeService is the concrete implementation of EmployeeService.
type employeeService struct {
	repo repository.EmployeeRepository
}

// NewEmployeeService creates a new instance of employeeService.
func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

// CreateEmployee creates a new employee through the repository.
func (s *employeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	return s.repo.Create(employee)
}

// GetEmployeeByID retrieves an employee by their ID using the repository.
func (s *employeeService) GetEmployeeByID(employeeID string) (*models.Employee, error) {
	return s.repo.FindByID(employeeID)
}

// GetAllEmployees retrieves all employees using the repository.
func (s *employeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.repo.FindAll()
}

// UpdateEmployee updates an existing employee using the repository.
func (s *employeeService) UpdateEmployee(employee *models.Employee) (*models.Employee, error) {
	return s.repo.Update(employee)
}

// DeleteEmployee deletes an employee by their ID using the repository.
func (s *employeeService) DeleteEmployee(employeeID string) error {
	return s.repo.Delete(employeeID)
}
