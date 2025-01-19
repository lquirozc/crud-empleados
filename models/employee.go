package models

import (
	"time"

	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
)

type Employee struct {
	EmployeeId mssql.UniqueIdentifier `gorm:"column:EmployeeId;primaryKey;type:uniqueidentifier" json:"employeeId"` // UUID como clave primaria
	IdNumber   string                 `gorm:"column:IdNumber;type:varchar(15);not null" json:"idNumber"`            // Documento de identidad
	Name       string                 `gorm:"column:Name;type:varchar(200);not null" json:"name"`                   // Nombre completo
	BornDate   time.Time              `gorm:"column:BornDate;type:datetime;not null" json:"bornDate"`               // Fecha de nacimiento
	SubAreaId  mssql.UniqueIdentifier `gorm:"column:SubAreaId;type:uniqueidentifier;not null" json:"subAreaId"`     // UUID de subárea
}

func (e *Employee) BeforeCreate(tx *gorm.DB) error {
	// Generar un nuevo UUID usando github.com/google/uuid
	newUUID := uuid.New()

	// Convertir el UUID generado a mssql.UniqueIdentifier
	e.EmployeeId = mssql.UniqueIdentifier(newUUID[:]) // Convertimos el UUID a []byte

	return nil
}

// Tabla personalizada en SQL Server
func (Employee) TableName() string {
	return "Employees"
}
