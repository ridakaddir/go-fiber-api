package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ridakaddir/go-fiber-api/database"
	"github.com/ridakaddir/go-fiber-api/models"
)

func CreateEmployee(c *fiber.Ctx) error {

	var employee models.Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&employee)

	return c.Status(201).JSON(employee)
}