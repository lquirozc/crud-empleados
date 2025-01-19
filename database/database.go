package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToSQLServer() error {

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error al cargar el archivo .env: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("No se pudo conectar a la base de datos: %v", err)
	}

	DB = db
	log.Println("Conexión a SQL Server exitosa")
	return nil
	// err = DB.AutoMigrate(&models.Employee{})
	// if err != nil {
	// 	log.Fatalf("Error al ejecutar las migraciones: %v", err)
	// 	return err
	// }

	// log.Println("Migración completada con éxito.")

}
