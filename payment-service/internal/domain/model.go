package domain

import "gorm.io/gorm"

type PaymentDetails struct {
	gorm.Model
	PaymentId string
	Amount    float64 `json:"amount" gorm:"not null"`
	OrderID   int     `json:"order_id" gorm:"not null"`
	Status    string
}
