package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	grpchandle "github.com/nuhmanudheent/hotel-booking/user-service/internal/delivery/grpc"
	httphandle "github.com/nuhmanudheent/hotel-booking/user-service/internal/delivery/http"
	"github.com/nuhmanudheent/hotel-booking/user-service/internal/repository"
	"github.com/nuhmanudheent/hotel-booking/user-service/internal/service"
	grpc_middle "github.com/nuhmanudheent/hotel-booking/user-service/pkg/middleware-grpc"
	"github.com/nuhmanudheent/hotel-booking/user-service/pkg/postgres"
	pb "github.com/nuhmanudheent/hotel-booking/user-service/proto"
	"google.golang.org/grpc"
)

var (
	port = ":50041"
)

func main() {
	conn, err := grpc.NewClient("localhost:50042", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	db := postgres.InitDatabase()
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo, conn)
	grpchandler := grpchandle.NewUserHandler(service)
	httpHandler := httphandle.NewUserHandler(service)

	secureMethods := map[string]bool{
		"/proto.UserService/UserGetInfo":   true,
		"/proto.UserService/GetHotelRooms": true,
		// Add other methods that require authentication
	}

	go func() {
		r := gin.Default()
		httpHandler.UserRouters(r)
		if err := r.Run(":8081"); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
		fmt.Println("gin server is running")
	}()

	go func() {
		list, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middle.AuthInterceptor(secureMethods)))
		pb.RegisterUserServiceServer(s, grpchandler)

		log.Printf("Server listening on port 50041")
		if err := s.Serve(list); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down servers...")
}
