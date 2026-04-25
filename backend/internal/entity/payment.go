package entity

import "time"

type Payment struct {
	ID           string
	MerchantName string
	Amount       float64
	Status       string
	CreatedAt    time.Time
}
