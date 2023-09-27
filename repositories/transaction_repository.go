package repositories

import (
	"context"
	"database/sql"
	"sanberhub/models/domain"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Find(ctx context.Context, tx *sql.Tx, savingAccountID int) ([]domain.Transaction, error)
}
