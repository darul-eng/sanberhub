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

	server.POST("/api/v1/daftar", userController.Register)
	server.POST("/api/v1/tabung", savingAccountController.Deposit)
	server.POST("/api/v1/tarik", savingAccountController.Withdraw)
	server.GET("/api/v1/saldo/:no_rekening", savingAccountController.Balance)
	server.GET("/api/v1/mutasi/:no_rekening", transactionController.Statement)

	if err := server.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
