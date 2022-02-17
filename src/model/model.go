package model

type (
	Funder struct {
		Id         uint64 `json:"id"`
		Address    string `json:"address"`
		Funded     bool   `json:"funded"`
		DateFunded string `json:"date"`
		Amount     uint64 `json:"amount"`
	}
)
