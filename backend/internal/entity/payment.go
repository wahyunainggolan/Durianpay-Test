package entity

import "time"

type Payment struct {
	ID           string
	MerchantName string
	Amount       float64
	Status       string
	CreatedAt    time.Time
}

type PaymentSummary struct {
	Total      int `json:"total"`
	Completed  int `json:"completed"`
	Processing int `json:"processing"`
	Failed     int `json:"failed"`
}
