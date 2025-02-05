package controllers

import (
	"net/http"

	"crud-empleados/models"
	"crud-empleados/services"

	"github.com/gin-gonic/gin"
)

// EmployeeController defines the methods the controller will expose.
type EmployeeController struct {
	service services.EmployeeService
}

// NewEmployeeController creates a new EmployeeController.
func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

// CreateEmployee handles the creation of a new employee.
func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Use the service to create the employee
	createdEmployee := ctrl.service.CreateEmployee(&employee)

	c.JSON(http.StatusCreated, createdEmployee)
}

// GetEmployeeByID retrieves an employee by their ID.
func (ctrl *EmployeeController) GetEmployeeByID(c *gin.Context) {
	employeeID := c.Param("id")

	// Use the service to get the employee
	employee := ctrl.service.GetEmployeeByID(employeeID)

	c.JSON(http.StatusOK, employee)
}

// GetAllEmployees retrieves all employees.
func (ctrl *EmployeeController) GetAllEmployees(c *gin.Context) {
	// Use the service to get all employees
	employees := ctrl.service.GetAllEmployees()

	c.JSON(http.StatusOK, employees)
}

// UpdateEmployee updates an existing employee.
func (ctrl *EmployeeController) UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Use the service to update the employee
	updatedEmployee := ctrl.service.UpdateEmployee(&employee)

	c.JSON(http.StatusOK, updatedEmployee)
}

// DeleteEmployee deletes an employee by their ID.
func (ctrl *EmployeeController) DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")

	// Use the service to delete the employee
	deletedEmployee := ctrl.service.DeleteEmployee(employeeID)

	c.JSON(http.StatusOK, deletedEmployee)
}
