package domain

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"nama"`
	NIK   int    `json:"nik"`
	Phone string `json:"no_hp"`
}
