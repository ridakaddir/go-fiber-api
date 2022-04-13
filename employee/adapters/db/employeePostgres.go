package db

import (
	"github.com/ridakaddir/go-fiber-api/database"
	"github.com/ridakaddir/go-fiber-api/employee/domain/models"
)

type EmployeePostgres struct {}

func NewEmployeePG () *EmployeePostgres {
	return &EmployeePostgres{}
}

func (employeePostgres *EmployeePostgres) CreateEmployee(employee *models.Employee) error{
	if database.Database.DB.Create(&employee).Error != nil {
		return database.Database.DB.Create(&employee).Error
	}
	return nil
}

func (employeePostgres *EmployeePostgres) GetEmployees() ([]models.Employee, error){
	var employees []models.Employee
	if database.Database.DB.Find(&employees).Error != nil {
		return nil, database.Database.DB.Find(&employees).Error
	}
	return employees, nil
}
