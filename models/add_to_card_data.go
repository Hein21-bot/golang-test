package models

import "time"

type AddToCard struct {
	User_id    string    `json:"user_id"`
	Product_id string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Added_at   time.Time `json:"added_at"`
}
