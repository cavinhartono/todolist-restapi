package models

type Product struct {
	ID int `json:"id"`
	Category string `json:"category"`
	Name string `json:"name"`
	Price int `json:"price"`
}
