package service

import (
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
	check := p.repo.PaymentConfirmation(c)
	if check{
		p.BookingService.BookingComplete()
	}
}
func (p *paymentService) PaymentCheck(orderId string) (string, error) {
	return p.repo.PaymentCheck(orderId)
}
func (p *paymentService) NewOrder(orderId string, price uint32) (string, error) {
	return p.repo.NewRazorOrder(orderId, price)
}
