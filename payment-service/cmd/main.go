package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	grpcHandler "github.com/nuhmanudheent/hotel-booking/payment-service/internal/delivery/grpc"
	httphandler "github.com/nuhmanudheent/hotel-booking/payment-service/internal/delivery/http"
	"github.com/nuhmanudheent/hotel-booking/payment-service/internal/repository"
	"github.com/nuhmanudheent/hotel-booking/payment-service/internal/service"
	"github.com/nuhmanudheent/hotel-booking/payment-service/pkg/postgres"
	pb "github.com/nuhmanudheent/hotel-booking/payment-service/proto"
	client_pb "github.com/nuhmanudheent/hotel-booking/payment-service/proto/client_proto"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("/home/nuhmanudheen-t/Desktop/go-microservice-hotel-booking/payment-service/.env")
	if err != nil {
		fmt.Println("Error to load env..............")
	}
	connBooking, err := grpc.NewClient("localhost:50043", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer connBooking.Close()
	userClient := client_pb.NewBookingServiceClient(connBooking)

	r := gin.Default()
	db := postgres.InitDatabase()
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo, userClient)
	paymentHandler := grpcHandler.NewPaymentHandler(paymentService)

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, paymentHandler)

	lis, err := net.Listen("tcp", ":50044")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	httpHandler := httphandler.NewPaymentHandler(paymentService)
	httpHandler.RegisterRoutes(r)

	// Load HTML templates
	r.LoadHTMLGlob("/home/nuhmanudheen-t/Desktop/go-microservice-hotel-booking/payment-service/templates/*")

	// Run the Gin server
	go func() {
		if err := r.Run(":8084"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Run the gRPC server
	go func() {
		log.Printf("gRPC server listening on :50044")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
}
