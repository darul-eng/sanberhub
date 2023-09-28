package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sanberhub/helpers"
	"sanberhub/models/api"
	"sanberhub/services"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Register(ctx echo.Context) error {
	registerRequest := api.RegisterRequest{}

	err := ctx.Bind(&registerRequest)
	helpers.PanicIfError(err)

	accountNumber, err := controller.UserService.Register(ctx, registerRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"remark": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":      http.StatusOK,
		"no_rekening": accountNumber,
	})
}
