package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ridakaddir/go-fiber-api/employee/domain/models"
	"github.com/ridakaddir/go-fiber-api/employee/domain/ports"
)

type HttpHandler struct {
	employeeService ports.EmployeeService
}

func EmployeeController (employeeService ports.EmployeeService) *HttpHandler {
	return &HttpHandler{employeeService: employeeService}
}

func (hdl *HttpHandler) CreateEmployee(c *fiber.Ctx) error {



	var employee models.Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	hdl.employeeService.Create(employee.FirstName, employee.LastName, employee.Email)

	return c.Status(201).JSON(employee)
}

func (hdl *HttpHandler) GetEmployees(c *fiber.Ctx) error {
	employees, err := hdl.employeeService.GetAll()
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(employees)
}