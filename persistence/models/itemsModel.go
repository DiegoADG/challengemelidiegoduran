package models

import (
	"time"
)

type ItemsModel struct {
	Site        string    `json:"site"`
	Id          string    `json:"id"`
	Price       float64   `json:"price"`
	Start_time  time.Time `json:"start_time"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Nickname    string    `json:"nickname"`
}
