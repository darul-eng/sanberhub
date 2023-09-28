package repositories

import (
	"context"
	"database/sql"
	"sanberhub/models/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Find(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
