package main

import (
	"clean-architecture-go/configs"
	"clean-architecture-go/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = configs.SetupDBConnection()
	authController controllers.AuthController = controllers.NewAuthController()
)

func main() {
	defer configs.CloseDatabaseConnection(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
