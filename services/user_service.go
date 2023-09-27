package services

import (
	"context"
	"sanberhub/models/api"
)

type UserService interface {
	Register(ctx context.Context, request api.RegisterRequest) int
}
