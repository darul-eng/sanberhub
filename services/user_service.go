package services

import (
	"github.com/labstack/echo/v4"
	"sanberhub/models/api"
)

type UserService interface {
	Register(ctx echo.Context, request api.RegisterRequest) (int, error)
}
