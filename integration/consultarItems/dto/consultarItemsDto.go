package dto

import "time"

type ConsultarItemsDto struct {
	Body itemsDto `json:"body"`
	Code int      `json:"code"`
}

type itemsDto struct {
	Price        float64   `json:"price"`
	Category_id  string    `json:"category_id"`
	Date_created time.Time `json:"date_created"`
	Seller_id    int       `json:"seller_id"`
	Currency_id  string    `json:"currency_id"`
}
