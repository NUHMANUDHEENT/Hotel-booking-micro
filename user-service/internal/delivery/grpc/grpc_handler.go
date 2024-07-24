package grpc

import (
	"context"
	"fmt"

	"github.com/nuhmanudheent/hotel-booking/user-service/internal/service"
	pb "github.com/nuhmanudheent/hotel-booking/user-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	err := u.service.RegisterUser(req.Name, req.Email, req.Password, req.Mobile)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterUserResponse{Message: "User registered successfully"}, nil
}

func (u *UserHandler) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := u.service.LoginUser(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginUserResponse{Message: "Login successful", JwtToken: user}, nil
}

func (u *UserHandler) UserGetInfo(ctx context.Context, req *pb.UserGetInfoRequest) (*pb.UserGetInfoResponse, error) {
	userID, ok := ctx.Value("userID").(uint)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	user, err := u.service.UserGetInfo(userID)
	if err != nil {
		return nil, err
	}
	return &pb.UserGetInfoResponse{User: &pb.User{
		Id:     uint32(user.ID),
		Name:   user.Name,
		Email:  user.Email,
		Mobile: user.Mobile,
	}}, nil
}
func (u *UserHandler) GetHotelRooms(ctx context.Context, req *pb.GetHotelRoomsRequest) (*pb.GetHotelRoomsResponse, error) {
	rooms, err := u.service.GetHotelRooms()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbRooms []*pb.Room
	for _, room := range rooms {
		pbRooms = append(pbRooms, &pb.Room{
			RoomNumber:   room.RoomNumber,
			Category:     room.Category,
			Availability: room.Availability,
			Price:        room.Price,
		})
	}

	return &pb.GetHotelRoomsResponse{Rooms: pbRooms}, nil
}
func (u *UserHandler) CheckUser(ctx context.Context, req *pb.CheckUserRequest) (*pb.CheckUserResponse, error) {
	fmt.Println("user check")
	exists := u.service.UserExists(req.UserId)
	return &pb.CheckUserResponse{Exists: exists}, nil
}
