package data

import "time"

type T struct {
	Users []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"users"`
	Orders []struct {
		ID          int       `json:"id"`
		UserID      int       `json:"userId"`
		OrderDate   time.Time `json:"orderDate"`
		TotalAmount float64   `json:"totalAmount"`
		Products    []int     `json:"products"`
	} `json:"orders"`
	Products []struct {
		ID    int     `json:"id"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	} `json:"products"`
}
