package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDatabase()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halo, ini server Go pertamaku!",
		})
	})

	r.GET("/projects", GetProjects)
	r.POST("/projects", CreateProject)
	r.PUT("/projects/:id", UpdateProject)
	r.DELETE("/projects/:id", DeleteProject)

	fmt.Println("Server jalan di http://localhost:8080")
	r.Run(":8080")
}
