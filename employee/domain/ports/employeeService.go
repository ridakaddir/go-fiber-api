package ports

import "github.com/ridakaddir/go-fiber-api/employee/domain/models"

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
	employee := models.Employee{FirstName: firstName, LastName: lastName, Email: email}
	return empolyeeService.employeeRepository.CreateEmployee(&employee)
}

func (empolyeeService *empolyeeService ) GetAll() ([]models.Employee, error) {
	return empolyeeService.employeeRepository.GetEmployees()
}
