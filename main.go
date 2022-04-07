package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ridakaddir/go-fiber-api/database"
	"github.com/ridakaddir/go-fiber-api/routes"
)

func welcome(c *fiber.Ctx) error  {
	return c.SendString("Welcome to Fiber")
}

func create(c *fiber.Ctx) error  {
	log.Println(string(c.Body()))
	log.Println(c.GetReqHeaders())
	return c.SendString("Create")
}

func main()  {

	database.ConnectDB()

	app:=fiber.New()
	
	app.Get("/api", welcome)

	app.Post("/api/employees", routes.CreateEmployee)

	port:= os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":"+port))
}