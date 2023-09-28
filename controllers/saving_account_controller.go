package controllers

import "github.com/labstack/echo/v4"

type SavingAccountController interface {
	Deposit(ctx echo.Context) error
	Withdraw(ctx echo.Context) error
	Balance(ctx echo.Context) error
}
