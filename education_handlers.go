package main

import (
	"github.com/gofiber/fiber/v2"
)

type Education struct {
	ID          int    `json:"id"`
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Period      string `json:"period"`
	Description string `json:"description"`
}

func GetEducation(c *fiber.Ctx) error {
	rows, err := DB.Query("SELECT id, institution, degree, period, description FROM education ORDER BY id")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var educations []Education

	for rows.Next() {
		var e Education
		err := rows.Scan(&e.ID, &e.Institution, &e.Degree, &e.Period, &e.Description)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		educations = append(educations, e)
	}

	return c.JSON(educations)
}

func CreateEducation(c *fiber.Ctx) error {
	var e Education

	if err := c.BodyParser(&e); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := DB.QueryRow(
		"INSERT INTO education (institution, degree, period, description) VALUES ($1, $2, $3, $4) RETURNING id",
		e.Institution, e.Degree, e.Period, e.Description,
	).Scan(&e.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(e)
}

func UpdateEducation(c *fiber.Ctx) error {
	id := c.Params("id")
	var e Education

	if err := c.BodyParser(&e); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := DB.Exec(
		"UPDATE education SET institution=$1, degree=$2, period=$3, description=$4 WHERE id=$5",
		e.Institution, e.Degree, e.Period, e.Description, id,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Education berhasil diupdate"})
}

func DeleteEducation(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := DB.Exec("DELETE FROM education WHERE id=$1", id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Education berhasil dihapus"})
}
