package services

import (
	"github.com/labstack/echo/v4"
	"sanberhub/models/api"
)

type SavingAccountService interface {
	Update(ctx echo.Context, request api.DepositOrWithdrawRequest) (float64, error)
	Find(ctx echo.Context, accountNumber int) (float64, error)
}
