package domain

import "gorm.io/gorm"

type PaymentDetails struct {
	gorm.Model
	PaymentId string
	Amount    float64 `json:"amount" `
	RazorOrderId   string     `json:"order_id"`
	Status    string
}
