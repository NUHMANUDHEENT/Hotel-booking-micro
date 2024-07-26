package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	grpcHandler "github.com/nuhmanudheent/hotel-booking/booking-service/internal/delivery/grpc"
	httpHandler "github.com/nuhmanudheent/hotel-booking/booking-service/internal/delivery/http"
	"github.com/nuhmanudheent/hotel-booking/booking-service/internal/repository"
	"github.com/nuhmanudheent/hotel-booking/booking-service/internal/service"
	"github.com/nuhmanudheent/hotel-booking/booking-service/pkg/postgres"
	pb "github.com/nuhmanudheent/hotel-booking/booking-service/proto"
	client_pb "github.com/nuhmanudheent/hotel-booking/booking-service/proto/client_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load("/home/nuhmanudheen-t/Desktop/go-microservice-hotel-booking/booking-service/.env")
	if err != nil {
		log.Println("No .env file found")
	}
	// Initialize the Gin router
	r := gin.Default()

	// Connect to other services
	connUser, err := grpc.NewClient("localhost:50041", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer connUser.Close()
	userClient := client_pb.NewUserServiceClient(connUser)

	connHotel, err := grpc.NewClient("localhost:50042", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to hotel service: %v", err)
	}
	defer connHotel.Close()
	hotelClient := client_pb.NewHotelServiceClient(connHotel)

	connPayment, err := grpc.NewClient("localhost:50044", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to payment service: %v", err)
	}
	paymentClient := client_pb.NewPaymentServiceClient(connPayment)
	defer connPayment.Close()
	db := postgres.InitDatabase()
	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepo, userClient, hotelClient, paymentClient)

	// Initialize the gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterBookingServiceServer(grpcServer, grpcHandler.NewBookingHandler(bookingService))

	// Set up the listener for gRPC server
	lis, err := net.Listen("tcp", ":50043")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	httpHandl := httpHandler.NewBookingHandler(bookingService)
	httpHandl.BookingRouters(r)

	// Run the Gin server
	go func() {
		if err := r.Run(":8083"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Run the gRPC server
	go func() {
		log.Printf("gRPC server listening on :50043")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down servers...")
}
