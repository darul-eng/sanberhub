package repositories

import (
	"context"
	"database/sql"
	"sanberhub/helpers"
	"sanberhub/models/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
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

func (repository *UserRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `SELECT "id", "name", "nik", "phone" FROM "user" WHERE "nik" = $1 OR "phone" = $2`
	rows, err := tx.QueryContext(ctx, SQL, user.NIK, user.Phone)
	helpers.PanicIfError(err)
	defer rows.Close()

	user = domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.NIK, &user.Phone)
		helpers.PanicIfError(err)
	}
	return user
}
