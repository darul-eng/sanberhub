package domain

type SavingAccount struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	AccountNumber int     `json:"nomor_rekening"`
	Balance       float64 `json:"saldo"`
}
