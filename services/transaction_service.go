package services

import (
	"github.com/labstack/echo/v4"
	"sanberhub/models/api"
)

type TransactionService interface {
	Find(ctx echo.Context, accountNumber int) ([]api.StatementResponse, error)
}
