package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/nuhmanudheent/hotel-booking/payment-service/internal/service"
	pb "github.com/nuhmanudheent/hotel-booking/payment-service/proto"
)

type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer
	Service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: service}
}
func (p *PaymentHandler) PaymentCheck(ctx context.Context, req *pb.PaymentCheckRequest) (*pb.PaymentCheckResponse, error) {
	status, err := p.Service.PaymentCheck(req.OrderId)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.PaymentCheckResponse{
		Status: status,
	}, nil
}
func (p *PaymentHandler) NewOrder(ctx context.Context, req *pb.NewOrderRequest) (*pb.NewOrderResponse, error) {
	fmt.Println("orde=====", req.OrderId)
	razorOrderId, err := p.Service.NewOrder(req.OrderId, req.Price)

	if err != nil {
		log.Fatal(err)
	}
	return &pb.NewOrderResponse{
		RazorOrderId: razorOrderId,
	}, nil
}
