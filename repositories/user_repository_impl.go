package repositories

import (
	"context"
	"database/sql"
	"sanberhub/helpers"
	"sanberhub/models/domain"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `INSERT INTO "user"("name", "nik", "phone") VALUES ($1, $2, $3) RETURNING id`
	var lastInsertID int
	result := tx.QueryRowContext(ctx, SQL, user.Name, user.NIK, user.Phone)

	err := result.Scan(&lastInsertID)
	helpers.PanicIfError(err)

	user.ID = lastInsertID

	return user
}
