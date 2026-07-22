package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ConnectDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Halo, ini server Go pertamaku!",
		})
	})

	app.Get("/projects", GetProjects)
	app.Post("/projects", AuthRequired, CreateProject)
	app.Put("/projects/:id", AuthRequired, UpdateProject)
	app.Delete("/projects/:id", AuthRequired, DeleteProject)

	app.Get("/education", GetEducation)
	app.Post("/education", AuthRequired, CreateEducation)
	app.Put("/education/:id", AuthRequired, UpdateEducation)
	app.Delete("/education/:id", AuthRequired, DeleteEducation)

	app.Post("/register", Register)
	app.Post("/login", Login)
	fmt.Println("Server jalan di http://localhost:8080")
	app.Listen(":8080")
}
