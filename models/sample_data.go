package models

import "time"

type ProductRequest struct {
	UPC           string                 `json:"upc"`
	AdminLevels   map[string]interface{} `json:"levels"`
	TimePurchased time.Time              `json:"time_purchased"`
	Name          string                 `json:"name"`
	Price         int                    `json:"price"`
	StockQuantity int                    `json:"stock_quantity"`
	Description   string                 `json:"description"`
}

// This will be the array structure for incoming requests
type BatchProductRequest []ProductRequest
