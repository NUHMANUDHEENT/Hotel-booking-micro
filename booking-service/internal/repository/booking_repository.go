package repository

import (
	"errors"

	"github.com/nuhmanudheent/hotel-booking/booking-service/internal/domain"
	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *domain.Booking) error
	UpdateBooking(booking *domain.Booking) error
	GetBookingByID(id uint) (*domain.Booking, error)
	BookingComplete(razorId string, status string) (string, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *domain.Booking) error {
	if err := r.db.Create(booking).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) UpdateBooking(booking *domain.Booking) error {
	if err := r.db.Save(booking).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) GetBookingByID(id uint) (*domain.Booking, error) {
	var booking domain.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &booking, nil
}
func (b *bookingRepository) BookingComplete(razorId string, status string) (string, error) {
	var booking domain.Booking
	if err := b.db.Where("razor_order_id = ?", razorId).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "Booking not found", err
		}
		return "", err
	}
	booking.Status = status
	if err := b.db.Save(&booking).Error; err != nil {
		return "", err
	}
	return "success", nil
}
