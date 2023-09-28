package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sanberhub/helpers"
	"sanberhub/services"
	"strconv"
)

type TransactionControllerImpl struct {
	TransactionService services.TransactionService
}

func NewTransactionController(TransactionService services.TransactionService) TransactionController {
	return &TransactionControllerImpl{TransactionService: TransactionService}
}

func (controller *TransactionControllerImpl) Statement(ctx echo.Context) error {
	param := ctx.Param("no_rekening")
	accountNumber, err := strconv.Atoi(param)
	helpers.PanicIfError(err)

	statement, err := controller.TransactionService.Find(ctx, accountNumber)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"remark": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"mutasi": statement,
	})

}
