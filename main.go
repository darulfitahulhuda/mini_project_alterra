package main

import (
	"main/controllers"
	"main/repository"
	"main/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	repository.Init()

	e := echo.New()
	e.Use(middleware.Logger())

	// Repository
	userRepository := repository.NewUserRepository(repository.DB)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepository)
	authUsecase := usecase.NewAuthUsecase()

	// Controllers
	userController := controllers.NewUserController(userUsecase, authUsecase)
	adminController := controllers.NewAdminController(userUsecase, authUsecase)

	// Auth routes
	e.POST("/register", userController.Create)
	e.POST("/login", userController.LoginUser)

	r := e.Group("/")
	{
		// jwt config
		jwtConfig := middleware.JWTConfig{
			SigningKey: []byte(usecase.SECRET_JWT),
		}
		//
		r.Use(middleware.JWTWithConfig(jwtConfig))

		// User routes
		r.GET("user", userController.GetUserByAuth)
		r.PUT("user", userController.UpdateUser)
		r.DELETE("user", userController.DeleteUser)

		// Admin routes
		r.GET("admin", adminController.GetAdminByAuth)
		r.GET("admin/users", adminController.GetAllUsers)
		r.PUT("admin", adminController.UpdateAdmin)
		r.DELETE("admin", adminController.DeleteAdmin)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
