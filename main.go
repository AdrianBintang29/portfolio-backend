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
	app.Post("/projects", CreateProject)
	app.Put("/projects/:id", UpdateProject)
	app.Delete("/projects/:id", DeleteProject)

	fmt.Println("Server jalan di http://localhost:8080")
	app.Listen(":8080")
}
