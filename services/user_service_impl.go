package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"sanberhub/helpers"
	"sanberhub/models/api"
	"sanberhub/models/domain"
	"sanberhub/repositories"
)

type UserServiceImpl struct {
	UserRepository          repositories.UserRepository
	SavingAccountRepository repositories.SavingAccount
	DB                      *sql.DB
	Validate                validator.Validate
}

func (service *UserServiceImpl) Register(ctx context.Context, request api.RegisterRequest) int {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user := domain.User{
		Name:  request.Name,
		NIK:   request.NIK,
		Phone: request.Phone,
	}

	user := service.UserRepository.Create(ctx, tx, user)

}
