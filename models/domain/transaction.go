package domain

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	SavingAccountID int       `json:"saving_account_id"`
	TransactionCode string    `json:"kode_transaksi"`
	Amount          float64   `json:"nominal"`
	CreatedAt       time.Time `json:"waktu"`
}
