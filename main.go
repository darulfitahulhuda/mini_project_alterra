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
	shoesRepository := repository.NewShoesRepository(repository.DB)
	transacationRepository := repository.NewTransactionRepository(repository.DB)
	cartRespository := repository.NewCartRepository(repository.DB)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepository)
	authUsecase := usecase.NewAuthUsecase()
	shoesUsecase := usecase.NewShoesUsecase(shoesRepository)
	transactionUsecase := usecase.NewTransactionUsecase(transacationRepository, shoesRepository)
	cartUsecase := usecase.NewCartUsecase(cartRespository, shoesRepository)

	// Controllers
	userController := controllers.NewUserController(userUsecase, authUsecase)
	adminController := controllers.NewAdminController(userUsecase, authUsecase)
	shoesController := controllers.NewShoesController(shoesUsecase, authUsecase)
	transactionController := controllers.NewTransactionController(transactionUsecase, authUsecase)
	cartController := controllers.NewCartContoller(cartUsecase, authUsecase)

	// Auth routes
	e.POST("/register", userController.Create)
	e.POST("/login", userController.LoginUser)

	r := e.Group("/")
	{
		// jwt config
		jwtConfig := middleware.JWTConfig{
			SigningKey: []byte(usecase.SECRET_JWT),
		}
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

		// shoes routes with auth admin
		r.POST("admin/shoes", shoesController.CreateShoes)
		r.PUT("admin/shoes/:id", shoesController.UpdateShoes)
		r.DELETE("admin/shoes/:id", shoesController.DeleteShoes)

		// Transaction routes
		r.POST("user/transaction", transactionController.CreateTransaction)
		r.GET("user/transaction", transactionController.GetTransactionByUser)
		r.GET("admin/transaction", transactionController.GetAllTransaction)
		r.PUT("admin/transaction/:id", transactionController.UpdateTransaction)
		r.PUT("admin/transaction/payment", transactionController.UpdatePaymentMethod)
		r.PATCH("admin/transaction/:id", transactionController.SoftDeleteTransaction)

		// Cart routes
		r.POST("user/carts", cartController.CreateCart)
		r.GET("user/carts", cartController.GetAllCarts)
		r.PUT("user/carts/:id", cartController.UpdateCart)
		r.DELETE("user/carts/:id", cartController.GetAllCarts)
	}

	// shoes routes
	e.GET("shoes", shoesController.GetAllShoes)
	e.GET("shoes/:id", shoesController.GetDetailShoes)

	e.Logger.Fatal(e.Start(":8080"))
}
