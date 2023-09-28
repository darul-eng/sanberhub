package services

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"math"
	"sanberhub/helpers"
	"sanberhub/models/api"
	"sanberhub/models/domain"
	"sanberhub/repositories"
	"time"
)

type SavingAccountServiceImpl struct {
	SavingAccountRepository repositories.SavingAccountRepository
	TransactionRepository   repositories.TransactionRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewSavingAccountService(savingAccount repositories.SavingAccountRepository, transactionRepository repositories.TransactionRepository, DB *sql.DB, validate *validator.Validate) SavingAccountService {
	return &SavingAccountServiceImpl{
		SavingAccountRepository: savingAccount,
		TransactionRepository:   transactionRepository,
		DB:                      DB,
		Validate:                validate,
	}
}

func (service *SavingAccountServiceImpl) Update(ctx echo.Context, request api.DepositOrWithdrawRequest) (float64, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return 0, err
	}

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	account, err := service.SavingAccountRepository.Find(ctx.Request().Context(), tx, request.AccountNumber)
	if err != nil {
		return 0, err
	}

	user := domain.SavingAccount{
		UserID:        1,
		AccountNumber: request.AccountNumber,
		Balance:       account.Balance + (request.Amount),
	}

	savingAccount := service.SavingAccountRepository.Update(ctx.Request().Context(), tx, user)

	var transactionCode string
	if request.Amount < 0 {
		transactionCode = "D"
	} else {
		transactionCode = "C"
	}

	transaction := domain.Transaction{
		UserID:          1,
		SavingAccountID: account.ID,
		TransactionCode: transactionCode,
		Amount:          math.Abs(request.Amount),
		CreatedAt:       time.Now(),
	}
	service.TransactionRepository.Create(ctx.Request().Context(), tx, transaction)

	return savingAccount.Balance, nil
}

func (service *SavingAccountServiceImpl) Find(ctx echo.Context, accountNumber int) (float64, error) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	account, err := service.SavingAccountRepository.Find(ctx.Request().Context(), tx, accountNumber)
	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}
