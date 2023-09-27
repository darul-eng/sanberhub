package domain

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	TransactionCode string    `json:"kode_transaksi"`
	Amount          float64   `json:"nominal"`
	CreatedAt       time.Time `json:"waktu"`
}
