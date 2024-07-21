package repository

import (
	"errors"

	"github.com/nuhmanudheent/hotel_booking/hotel_service/internal/domain"
	"gorm.io/gorm"
)

type RoomRepository interface {
	AddRoom(room domain.Room) error
	GetRoom(roomNumber string) (domain.Room, error)
	UpdateRoom(room domain.Room) error
	DeleteRoom(roomNumber string) error
	GetRooms() ([]domain.Room, error) // Add this line
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) AddRoom(room domain.Room) error {
	if err := r.db.Create(&room).Error; err != nil {
		return errors.New("failed to add room")
	}
	return nil
}

func (r *roomRepository) GetRoom(roomNumber string) (domain.Room, error) {
	var room domain.Room
	if err := r.db.First(&room, "room_number = ?", roomNumber).Error; err != nil {
		return domain.Room{}, errors.New("room not found")
	}
	return room, nil
}

func (r *roomRepository) UpdateRoom(room domain.Room) error {
	if err := r.db.Where("room_number = ?", room.RoomNumber).Save(&room).Error; err != nil {
		return errors.New("failed to update room")
	}
	return nil
}

func (r *roomRepository) DeleteRoom(roomNumber string) error {
	if err := r.db.Delete(&domain.Room{}, "room_number = ?", roomNumber).Error; err != nil {
		return errors.New("failed to delete room")
	}
	return nil
}

func (r *roomRepository) GetRooms() ([]domain.Room, error) { // Add this function
	var rooms []domain.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, errors.New("failed to get rooms")
	}
	return rooms, nil
}
