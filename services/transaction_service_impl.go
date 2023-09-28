package services

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"sanberhub/helpers"
	"sanberhub/models/api"
	"sanberhub/repositories"
)

type TransactionServiceImpl struct {
	SavingAccountRepository repositories.SavingAccountRepository
	TransactionRepository   repositories.TransactionRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewTransactionService(savingAccountRepository repositories.SavingAccountRepository, TransactionRepository repositories.TransactionRepository, DB *sql.DB, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		SavingAccountRepository: savingAccountRepository,
		TransactionRepository:   TransactionRepository,
		DB:                      DB,
		Validate:                validate,
	}
}

func (service *TransactionServiceImpl) Find(ctx echo.Context, accountNumber int) ([]api.StatementResponse, error) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	account, err := service.SavingAccountRepository.Find(ctx.Request().Context(), tx, accountNumber)
	if err != nil {
		return nil, err
	}

	transactions, err := service.TransactionRepository.Find(ctx.Request().Context(), tx, account.ID)
	if err != nil {
		return nil, err
	}

	return helpers.ToStatementResponses(transactions), nil
}
