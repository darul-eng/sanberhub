package repositories

import (
	"context"
	"database/sql"
	"errors"
	"sanberhub/helpers"
	"sanberhub/models/domain"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `INSERT INTO "transaction"("user_id", "saving_account_id", "transaction_code", "amount", "created_at") VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, transaction.UserID, transaction.SavingAccountID, transaction.TransactionCode, transaction.Amount, transaction.CreatedAt)

	err := result.Scan(&lastInsertId)
	helpers.PanicIfError(err)

	transaction.ID = lastInsertId

	return transaction
}

func (repository *TransactionRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, savingAccountID int) ([]domain.Transaction, error) {
	SQL := `SELECT "id", "user_id", "saving_account_id", "transaction_code", "amount", "created_at" FROM "transaction" WHERE "saving_account_id" = $1`
	rows, err := tx.QueryContext(ctx, SQL, savingAccountID)
	helpers.PanicIfError(err)
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.SavingAccountID, &transaction.TransactionCode, &transaction.Amount, &transaction.CreatedAt)
		helpers.PanicIfError(err)

		transactions = append(transactions, transaction)
	}

	if len(transactions) == 0 {
		return transactions, errors.New("transaction is not found")
	}

	return transactions, nil
}
