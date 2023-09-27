package repositories

import (
	"context"
	"database/sql"
	"errors"
	"sanberhub/helpers"
	"sanberhub/models/domain"
)

type SavingAccountImpl struct {
}

func (repository *SavingAccountImpl) Create(ctx context.Context, tx *sql.Tx, account domain.SavingAccount) domain.SavingAccount {
	SQL := `INSERT INTO "saving_account"("user_id", "account_number", "balance") VALUES ($1, $2, $3) RETURNING id`
	var lastInsertID int
	result := tx.QueryRowContext(ctx, SQL, account.UserID, account.AccountNumber, account.Balance)

	err := result.Scan(&lastInsertID)
	helpers.PanicIfError(err)

	account.ID = lastInsertID

	return account
}

func (repository *SavingAccountImpl) Update(ctx context.Context, tx *sql.Tx, account domain.SavingAccount) domain.SavingAccount {
	SQL := `UPDATE "saving_account" SET "balance"=$1 WHERE "account_number"=$2 AND "user_id"=$3`
	_, err := tx.ExecContext(ctx, SQL, account.Balance, account.AccountNumber, account.UserID)
	helpers.PanicIfError(err)

	return account
}

func (repository *SavingAccountImpl) Find(ctx context.Context, tx *sql.Tx, accountNumber int) (domain.SavingAccount, error) {
	SQL := `SELECT "user_id", "account_number", "balance" FROM "saving_account" WHERE "account_number" = $1`
	rows, err := tx.QueryContext(ctx, SQL, accountNumber)
	helpers.PanicIfError(err)
	defer rows.Close()

	savingAccount := domain.SavingAccount{}
	if rows.Next() {
		err := rows.Scan(&savingAccount.UserID, &savingAccount.AccountNumber, &savingAccount.Balance)
		helpers.PanicIfError(err)

		return savingAccount, nil
	} else {
		return savingAccount, errors.New("account is not found")
	}
}
