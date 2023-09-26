package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"sanberhub/config"
	"sanberhub/helpers"
)

func main() {
	err := godotenv.Load()
	helpers.PanicIfError(err)

	server := echo.New()
	config.NewDB()

	if err := server.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
