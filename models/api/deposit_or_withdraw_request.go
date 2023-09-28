package api

type DepositOrWithdrawRequest struct {
	AccountNumber int     `validate:"required"json:"no_rekening"`
	Amount        float64 `validate:"required"json:"nominal"`
}
