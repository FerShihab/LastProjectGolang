package models

type Car struct {
	Id        int    `json:"id"`
	Merk      string `json:"merk"`
	LuxuryCar bool   `json:"luxurycar"`
}
