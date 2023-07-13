package main

import (
	"github.com/cleanupDev/verbose-pancake/backend/internal/repositories"
	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", helloWorld)
	router.GET("/ping", repositories.PingDatabase)
	router.GET("/initdb", repositories.InitDB)
	router.GET("/show/users", repositories.ShowUsers)
	router.POST("/create/user", repositories.CreateUser)

	router.Run("localhost:5001")
}
