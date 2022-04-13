package ports

import (
	"errors"
	"log"

	"github.com/ridakaddir/go-fiber-api/employee/domain/models"
)

type EmployeeService interface {
	Create( firstName string, lastName string,  email string ) error
	GetAll() ([]models.Employee, error)
}

type empolyeeService struct {
	employeeRepository employeeRepository
}

func NewEmployeeService (employeeRepository employeeRepository) *empolyeeService {
	return &empolyeeService{employeeRepository: employeeRepository}
}

func (empolyeeService *empolyeeService ) Create( firstName string, lastName string,  email string ) error {
	if firstName == "" || lastName == "" || email == "" {
		log.Println("Error: FirstName, LastName and Email are required")
		return errors.New("First name, last name and email are required")
	}
	employee := models.Employee{FirstName: firstName, LastName: lastName, Email: email}
	 empolyeeService.employeeRepository.CreateEmployee(&employee)
	 return nil
}

func (empolyeeService *empolyeeService ) GetAll() ([]models.Employee, error) {
	return empolyeeService.employeeRepository.GetEmployees()
}
