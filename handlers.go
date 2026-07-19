package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

func GetProjects(c *gin.Context) {
	rows, err := DB.Query("SELECT id, title, description, link FROM projects ORDER BY id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var p Project
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Link)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		projects = append(projects, p)
	}

	c.JSON(http.StatusOK, projects)
}

func CreateProject(c *gin.Context) {
	var p Project

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := DB.QueryRow(
		"INSERT INTO projects (title, description, link) VALUES ($1, $2, $3) RETURNING id",
		p.Title, p.Description, p.Link,
	).Scan(&p.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var p Project

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec(
		"UPDATE projects SET title=$1, description=$2, link=$3 WHERE id=$4",
		p.Title, p.Description, p.Link, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project berhasil diupdate"})
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	_, err := DB.Exec("DELETE FROM projects WHERE id=$1", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project berhasil dihapus"})
}
