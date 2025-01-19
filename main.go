package main

import (
	"log"

	"crud-empleados/database"
	"crud-empleados/routes"
)

func main() {
	// Conectar a la base de datos
	err := database.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return
	}

	// Configurar las rutas de la API
	router := routes.ConfigurarRutas()

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
