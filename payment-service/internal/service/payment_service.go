package service

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nuhmanudheent/hotel-booking/payment-service/internal/repository"
	client_pb "github.com/nuhmanudheent/hotel-booking/payment-service/proto/client_proto"
)

type PaymentService interface {
	PaymentComplete(c *gin.Context)
	PaymentConfirmation(c *gin.Context)
	PaymentCheck(orderId string) (string, error)
	NewOrder(orderId string, price uint32) (string, error)
}

type paymentService struct {
	repo           repository.PaymentRepository
	BookingService client_pb.BookingServiceClient
}

func NewPaymentService(repo repository.PaymentRepository, bookingservice client_pb.BookingServiceClient) PaymentService {
	return &paymentService{repo: repo, BookingService: bookingservice}
}

func (p *paymentService) PaymentComplete(c *gin.Context) {
	p.repo.PaymentComplete(c)
}
func (p *paymentService) PaymentConfirmation(c *gin.Context) {
	razorId, status := p.repo.PaymentConfirmation(c)
	if status != "success" {
		log.Fatal("Payment failed")
	}
	resp, err := p.BookingService.BookingComplete(context.Background(), &client_pb.BookingCompleteRequest{
		OrderId: razorId,
		Status:  status,
	})
	if err != nil {
		log.Fatal("failed to update booking data")
	}
	fmt.Println("Payment ccomplete", resp.Status)
}
func (p *paymentService) PaymentCheck(orderId string) (string, error) {
	return p.repo.PaymentCheck(orderId)
}
func (p *paymentService) NewOrder(orderId string, price uint32) (string, error) {
	return p.repo.NewRazorOrder(orderId, price)
}
