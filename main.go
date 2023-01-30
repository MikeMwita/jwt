package main

import (
	routes "github.com/MikeMwita/JWT_project/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api1"})
	})

	router.GET("/api2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api2"})
	})
	router.Run(":" + port)
}
