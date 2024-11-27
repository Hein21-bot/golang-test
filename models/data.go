package models

import "time"

type Data struct {
	UPC           string
	AdminLevels   map[string]interface{}
	TimePurchased []time.Time
}
