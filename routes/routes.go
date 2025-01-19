package routes

import (
	"crud-empleados/controllers"
	"crud-empleados/database"
	repository "crud-empleados/repositories"
	"crud-empleados/services"

	"github.com/gin-gonic/gin"
)

// ConfigurarRutas establece todas las rutas de la API y las dependencias.
func ConfigurarRutas() *gin.Engine {
	// Crear la conexión a la base de datos
	//database.ConnectToSQLServer()

	// Crear repositorio e inyectarlo en el servicio
	employeeRepo := repository.NewEmployeeRepository(database.DB)
	employeeService := services.NewEmployeeService(employeeRepo)

	// Crear el controlador e inyectar el servicio
	employeeController := controllers.NewEmployeeController(employeeService)

	// Crear un router Gin
	router := gin.Default()

	// Definir las rutas y los métodos HTTP
	router.POST("/employees", employeeController.CreateEmployee)
	router.GET("/employees/:id", employeeController.GetEmployeeByID)
	router.GET("/employees", employeeController.GetAllEmployees)
	router.PUT("/employees", employeeController.UpdateEmployee)
	router.DELETE("/employees/:id", employeeController.DeleteEmployee)

	return router
}
