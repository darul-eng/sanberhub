package repositories

import (
	"context"
	"database/sql"
	"sanberhub/models/domain"
)

type SavingAccount interface {
	Create(ctx context.Context, tx *sql.Tx, account domain.SavingAccount) domain.SavingAccount
	Update(ctx context.Context, tx *sql.Tx, account domain.SavingAccount) domain.SavingAccount
	Find(ctx context.Context, tx *sql.Tx, accountNumber int) (domain.SavingAccount, error)
}
