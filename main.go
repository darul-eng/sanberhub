package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"sanberhub/config"
	"sanberhub/controllers"
	"sanberhub/helpers"
	"sanberhub/repositories"
	"sanberhub/services"
)

func main() {
	err := godotenv.Load()
	helpers.PanicIfError(err)

	server := echo.New()
	db := config.NewDB()
	validate := validator.New()

	userRepository := repositories.NewUserRepository()
	savingAccountRepository := repositories.NewSavingAccountRepository()
	userService := services.NewUserService(userRepository, savingAccountRepository, db, validate)
	userController := controllers.NewUserController(userService)

	transactionRepository := repositories.NewTransactionRepository()
	savingAccountService := services.NewSavingAccountService(savingAccountRepository, transactionRepository, db, validate)
	savingAccountController := controllers.NewSavingAccountController(savingAccountService)

	transactionService := services.NewTransactionService(savingAccountRepository, transactionRepository, db, validate)
	transactionController := controllers.NewTransactionController(transactionService)

	routeGroup := server.Group("/api/v1")
	routeGroup.POST("/daftar", userController.Register)
	routeGroup.POST("/tabung", savingAccountController.Deposit)
	routeGroup.POST("/tarik", savingAccountController.Withdraw)
	routeGroup.GET("/saldo/:no_rekening", savingAccountController.Balance)
	routeGroup.GET("/mutasi/:no_rekening", transactionController.Statement)

	if err := server.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
