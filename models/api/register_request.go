package api

type RegisterRequest struct {
	Name  string `validate:"required"json:"nama"`
	NIK   int    `validate:"required"json:"nik"`
	Phone string `validate:"required"json:"no_hp"`
}
