package main

import (
	"os"

	"github.com/cleanupDev/verbose-pancake/backend/internal/handlers"
	"github.com/cleanupDev/verbose-pancake/backend/internal/repositories"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}

	router.Use(cors.New(config))

	router.GET("/", helloWorld)
	router.GET("/ping", repositories.PingDatabase)
	router.GET("/initdb", repositories.InitDB)
	router.GET("/show/users", handlers.ShowUsers)
	router.POST("/create/user", handlers.CreateUser)
	router.POST("/login", handlers.LoginUser)
	router.POST("/user", handlers.GetUser)

	router.Run(os.Getenv("BACKEND_URL"))
}
