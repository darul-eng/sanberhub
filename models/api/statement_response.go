package api

import "time"

type StatementResponse struct {
	CreatedAt       time.Time `json:"waktu"`
	TransactionCode string    `json:"kode_transaksi"`
	Amount          float64   `json:"nominal"`
}
