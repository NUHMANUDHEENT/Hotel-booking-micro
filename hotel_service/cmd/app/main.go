package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	grpcHandler "github.com/nuhmanudheent/hotel_booking/hotel_service/internal/delivery/grpc"
	httpHandler "github.com/nuhmanudheent/hotel_booking/hotel_service/internal/delivery/http"
	"github.com/nuhmanudheent/hotel_booking/hotel_service/internal/repository"
	"github.com/nuhmanudheent/hotel_booking/hotel_service/internal/service"
	"github.com/nuhmanudheent/hotel_booking/hotel_service/pkg/postgres"
	pb "github.com/nuhmanudheent/hotel_booking/hotel_service/proto"
)

func main() {

	db := postgres.InitDatabase()
	roomRepo := repository.NewRoomRepository(db)
	hotelService := service.NewHotelService(roomRepo)

	r := gin.Default()
	httpHandler := httpHandler.NewHotelHandler(hotelService)
	httpHandler.HotelRouters(r)

	grpcServer := grpc.NewServer()
	grpcHandler := grpcHandler.NewHotelHandler(hotelService)
	pb.RegisterHotelServiceServer(grpcServer, grpcHandler)

	go func() {
		if err := r.Run(":8082"); err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50042")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server is running on port 50042")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
