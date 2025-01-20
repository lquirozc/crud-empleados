package services

import (
	"crud-empleados/models"
	repository "crud-empleados/repositories"
)

// EmployeeService defines the methods that the service layer will provide.
type EmployeeService interface {
	CreateEmployee(employee *models.Employee) models.Result[models.Employee]
	GetEmployeeByID(employeeID string) models.Result[models.Employee]
	GetAllEmployees() models.Result[[]models.Employee]
	UpdateEmployee(employee *models.Employee) models.Result[models.Employee]
	DeleteEmployee(employeeID string) models.Result[string]
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
func (s *employeeService) CreateEmployee(employee *models.Employee) models.Result[models.Employee] {

	createdEmployee, err := s.repo.Create(employee)
	if err != nil {
		return models.Failure[models.Employee]("Cannot create employee: " + err.Error())
	}

	return models.Success(*createdEmployee)
}

// GetEmployeeByID retrieves an employee by their ID using the repository.
func (s *employeeService) GetEmployeeByID(employeeID string) models.Result[models.Employee] {

	findedEmployee, err := s.repo.FindByID(employeeID)

	if err != nil {
		return models.Failure[models.Employee]("Employee not found: " + err.Error())
	}

	return models.Success(*findedEmployee)

}

// GetAllEmployees retrieves all employees using the repository.
func (s *employeeService) GetAllEmployees() models.Result[[]models.Employee] {

	employees, err := s.repo.FindAll()
	if err != nil {
		return models.Failure[[]models.Employee]("Error retrieving employees: " + err.Error())
	}

	return models.Success(employees)
}

// UpdateEmployee updates an existing employee using the repository.
func (s *employeeService) UpdateEmployee(employee *models.Employee) models.Result[models.Employee] {

	updatedEmployee, err := s.repo.Update(employee)
	if err != nil {
		return models.Failure[models.Employee]("Error updating employee: " + err.Error())
	}

	return models.Success(*updatedEmployee)
}

// DeleteEmployee deletes an employee by their ID using the repository.
func (s *employeeService) DeleteEmployee(employeeID string) models.Result[string] {

	if err := s.repo.Delete(employeeID); err != nil {
		return models.Failure[string]("Error deleting employee: " + err.Error())
	}

	return models.Success("Employee deleted successfully")
}
