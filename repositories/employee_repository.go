package repository

import (
	"crud-empleados/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) (*models.Employee, error)
	FindByID(employeeID string) (*models.Employee, error)
	FindAll() ([]models.Employee, error)
	Update(employee *models.Employee) (*models.Employee, error)
	Delete(employeeID string) error
}

// employeeRepository is the concrete implementation of EmployeeRepository.
type employeeRepository struct {
	db *gorm.DB
}

// NewEmployeeRepository creates a new instance of employeeRepository.
func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

// Create a new employee in the database.
func (r *employeeRepository) Create(employee *models.Employee) (*models.Employee, error) {
	if err := r.db.Create(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

// FindByID retrieves an employee by their ID.
func (r *employeeRepository) FindByID(employeeID string) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.First(&employee, "employeeId = ?", employeeID).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

// FindAll retrieves all employees from the database.
func (r *employeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	if err := r.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// Update modifies an existing employee in the database.
func (r *employeeRepository) Update(employee *models.Employee) (*models.Employee, error) {
	if err := r.db.Save(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

// Delete removes an employee from the database by their ID.
func (r *employeeRepository) Delete(employeeID string) error {
	if err := r.db.Delete(&models.Employee{}, "employeeId = ?", employeeID).Error; err != nil {
		return err
	}
	return nil
}
