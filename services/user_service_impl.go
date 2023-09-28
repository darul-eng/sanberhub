package services

import (
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"math/rand"
	"sanberhub/helpers"
	"sanberhub/models/api"
	"sanberhub/models/domain"
	"sanberhub/repositories"
)

type UserServiceImpl struct {
	UserRepository          repositories.UserRepository
	SavingAccountRepository repositories.SavingAccountRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewUserService(userRepository repositories.UserRepository, savingAccountRepository repositories.SavingAccountRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository:          userRepository,
		SavingAccountRepository: savingAccountRepository,
		DB:                      DB,
		Validate:                validate,
	}
}

func (service *UserServiceImpl) Register(ctx echo.Context, request api.RegisterRequest) (int, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return 0, err
	}

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user := domain.User{
		Name:  request.Name,
		NIK:   request.NIK,
		Phone: request.Phone,
	}

	userFind := service.UserRepository.Find(ctx.Request().Context(), tx, user)
	if userFind.ID != 0 {
		return 0, errors.New("nik atau nomor hp telah digunakan")
	}

	userResponse := service.UserRepository.Create(ctx.Request().Context(), tx, user)

	accountNumber := rand.Intn(userResponse.ID) + 110201
	savingAccount := domain.SavingAccount{
		UserID:        userResponse.ID,
		AccountNumber: accountNumber,
		Balance:       100000,
	}

	savingAccountResponse := service.SavingAccountRepository.Create(ctx.Request().Context(), tx, savingAccount)

	return savingAccountResponse.AccountNumber, nil
}
