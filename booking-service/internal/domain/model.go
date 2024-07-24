package domain

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	OrderId      string
	UserID       uint
	RoomID       string `json:"room_id"`
	CheckIn      time.Time
	CheckOut     time.Time
	Amount       float64 `json:"amount"`
	Status       string
	RazorOrderId string
}
