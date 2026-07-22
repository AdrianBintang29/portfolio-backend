package main

import (
	"github.com/gofiber/fiber/v2"
)

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

func GetProjects(c *fiber.Ctx) error {
	rows, err := DB.Query("SELECT id, title, description, link FROM projects ORDER BY id")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var p Project
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Link)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		projects = append(projects, p)
	}

	return c.JSON(projects)
}

func CreateProject(c *fiber.Ctx) error {
	var p Project

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := DB.QueryRow(
		"INSERT INTO projects (title, description, link) VALUES ($1, $2, $3) RETURNING id",
		p.Title, p.Description, p.Link,
	).Scan(&p.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(p)
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var p Project

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := DB.Exec(
		"UPDATE projects SET title=$1, description=$2, link=$3 WHERE id=$4",
		p.Title, p.Description, p.Link, id,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Project berhasil diupdate"})
}

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := DB.Exec("DELETE FROM projects WHERE id=$1", id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Project berhasil dihapus"})
}
