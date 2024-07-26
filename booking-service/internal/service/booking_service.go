package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nuhmanudheent/hotel-booking/booking-service/internal/domain"
	"github.com/nuhmanudheent/hotel-booking/booking-service/internal/repository"
	client_pb "github.com/nuhmanudheent/hotel-booking/booking-service/proto/client_proto"
)

type BookingService interface {
	CreateBooking(userID uint32, roomID string, checkIn, checkOut time.Time, amount float64) (*domain.Booking, error)
	BookingComplete(razorId string, status string) (bool, error)
}

type bookingService struct {
	repo           repository.BookingRepository
	userService    client_pb.UserServiceClient
	hotelService   client_pb.HotelServiceClient
	paymentService client_pb.PaymentServiceClient
}

func NewBookingService(repo repository.BookingRepository, userService client_pb.UserServiceClient, hotelService client_pb.HotelServiceClient, paymentService client_pb.PaymentServiceClient) BookingService {
	return &bookingService{
		repo:           repo,
		userService:    userService,
		hotelService:   hotelService,
		paymentService: paymentService,
	}
}

func (s *bookingService) CreateBooking(userID uint32, roomID string, checkIn, checkOut time.Time, amount float64) (*domain.Booking, error) {
	user, err := s.userService.CheckUser(context.Background(), &client_pb.CheckUserRequest{UserId: userID})
	if err != nil || !user.Exists {
		return nil, fmt.Errorf("user is not found")
	}
	fmt.Println("user--------", userID)
	room, err := s.hotelService.CheckRoom(context.Background(), &client_pb.CheckRoomRequest{RoomId: roomID})
	if err != nil || !room.Available {
		return nil, fmt.Errorf("room is not available: %v", err)
	}

	orderID := uuid.New().String()[:5]
	// Process payment
	paymentResp, err := s.paymentService.NewOrder(context.Background(), &client_pb.NewOrderRequest{
		OrderId: orderID,
		Price:   uint32(amount),
	})
	if err != nil || paymentResp.RazorOrderId == "" {
		return nil, fmt.Errorf("orderID processing failed: %v", err)
	}
	// Create booking
	booking := &domain.Booking{
		UserID:       uint(userID),
		RoomID:       roomID,
		CheckIn:      checkIn,
		CheckOut:     checkOut,
		Amount:       amount,
		Status:       "Pending",
		OrderId:      orderID,
		RazorOrderId: paymentResp.RazorOrderId,
	}

	if err := s.repo.CreateBooking(booking); err != nil {
		return nil, fmt.Errorf("failed to create booking: %v", err)
	}

	return booking, nil
}
func (b *bookingService) BookingComplete(razorId string, status string) (bool, error) {
	// Process payment
	status, err := b.repo.BookingComplete(razorId, status)
	if err != nil {
		return false, fmt.Errorf("failed to complete booking: %v", err)
	}
	return true, nil

}
