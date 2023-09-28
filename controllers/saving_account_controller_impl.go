package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sanberhub/helpers"
	"sanberhub/models/api"
	"sanberhub/services"
	"strconv"
)

type SavingAccountControllerImpl struct {
	SavingAccountService services.SavingAccountService
}

func NewSavingAccountController(service services.SavingAccountService) SavingAccountController {
	return &SavingAccountControllerImpl{SavingAccountService: service}
}

func (controller *SavingAccountControllerImpl) Deposit(ctx echo.Context) error {
	depositRequest := api.DepositOrWithdrawRequest{}

	err := ctx.Bind(&depositRequest)
	helpers.PanicIfError(err)

	balance, err := controller.SavingAccountService.Update(ctx, depositRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"remark": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"saldo":  balance,
	})
}

func (controller *SavingAccountControllerImpl) Withdraw(ctx echo.Context) error {
	depositRequest := api.DepositOrWithdrawRequest{}

	err := ctx.Bind(&depositRequest)
	helpers.PanicIfError(err)

	depositRequest.Amount = -depositRequest.Amount
	balance, err := controller.SavingAccountService.Update(ctx, depositRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"remark": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"saldo":  balance,
	})
}

func (controller *SavingAccountControllerImpl) Balance(ctx echo.Context) error {
	param := ctx.Param("no_rekening")
	accountNumber, err := strconv.Atoi(param)
	helpers.PanicIfError(err)

	balance, err := controller.SavingAccountService.Find(ctx, accountNumber)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"remark": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"saldo":  balance,
	})

}
