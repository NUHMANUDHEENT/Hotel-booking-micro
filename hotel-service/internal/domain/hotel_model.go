package domain

type Room struct {
	RoomNumber   string  `gorm:"unique" ,json:"roomnumber"`
	Category     string  `json:"category"`
	Availability bool    `json:"availability"`
	Price        float64 `json:"price"`
}
