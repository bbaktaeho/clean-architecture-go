package main

import (
	"clean-architecture-go/configs"
	"clean-architecture-go/controllers"
	"clean-architecture-go/repositories"
	"clean-architecture-go/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = configs.SetupDBConnection()
	userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	jwtService     services.JWTService         = services.NewJWTService()
	authService    services.AuthService        = services.NewAuthService(userRepository)
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
)

func main() {
	defer configs.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
