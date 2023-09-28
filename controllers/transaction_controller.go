package controllers

import "github.com/labstack/echo/v4"

type TransactionController interface {
	Statement(ctx echo.Context) error
}
