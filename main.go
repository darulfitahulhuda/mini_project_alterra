package main

import (
	"main/controllers"
	"main/repository"
	"main/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	repository.Init()

	e := echo.New()

	// Repository
	userRepository := repository.NewUserRepository(repository.DB)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepository)
	authUsecase := usecase.NewAuthUsecase()

	// Controllers
	userController := controllers.NewUserController(userUsecase, authUsecase)

	e.POST("/users", userController.Create)
	e.POST("/login", userController.LoginUser)
	e.GET("/users", userController.GetAllUsers)

	e.Logger.Fatal(e.Start(":8080"))

}
