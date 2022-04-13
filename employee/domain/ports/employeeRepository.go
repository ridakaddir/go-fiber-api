package ports

import "github.com/ridakaddir/go-fiber-api/employee/domain/models"

type employeeRepository interface {
	CreateEmployee(employee *models.Employee) error
	GetEmployees() ([]models.Employee, error)
}