package service

import (
	"github.com/nuhmanudheent/hotel_booking/hotel_service/internal/domain"
	"github.com/nuhmanudheent/hotel_booking/hotel_service/internal/repository"
)

type HotelService interface {
	AddRoom(room domain.Room) error
	GetRoom(roomNumber string) (domain.Room, error)
	UpdateRoom(room domain.Room) error
	DeleteRoom(roomNumber string) error
	GetRooms() ([]domain.Room, error)
	RoomAvailable(roomID string) bool
}

type hotelService struct {
	repo repository.RoomRepository
}

func NewHotelService(repo repository.RoomRepository) HotelService {
	return &hotelService{repo: repo}
}

func (s *hotelService) AddRoom(room domain.Room) error {
	return s.repo.AddRoom(room)
}

func (s *hotelService) GetRoom(roomNumber string) (domain.Room, error) {
	return s.repo.GetRoom(roomNumber)
}

func (s *hotelService) UpdateRoom(room domain.Room) error {
	return s.repo.UpdateRoom(room)
}

func (s *hotelService) DeleteRoom(roomNumber string) error {
	return s.repo.DeleteRoom(roomNumber)
}

func (s *hotelService) GetRooms() ([]domain.Room, error) { // Implement this function
	return s.repo.GetRooms()
}
func (s *hotelService) RoomAvailable(roomID string) bool {
	return s.repo.RoomAvailable(roomID)
}
