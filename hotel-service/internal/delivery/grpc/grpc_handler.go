package grpc

import (
	"context"

	"github.com/nuhmanudheent/hotel_booking/hotel-service/internal/domain"
	"github.com/nuhmanudheent/hotel_booking/hotel-service/internal/service"
	pb "github.com/nuhmanudheent/hotel_booking/hotel-service/proto"
)

type HotelHandler struct {
	pb.UnimplementedHotelServiceServer
	service service.HotelService
}

func NewHotelHandler(service service.HotelService) *HotelHandler {
	return &HotelHandler{service: service}
}

func (h *HotelHandler) AddRoom(ctx context.Context, req *pb.AddRoomRequest) (*pb.AddRoomResponse, error) {
	room := domain.Room{
		RoomNumber:   req.Room.RoomNumber,
		Category:     req.Room.Category,
		Availability: req.Room.Availability,
		Price:        req.Room.Price,
	}
	err := h.service.AddRoom(room)
	if err != nil {
		return nil, err
	}
	return &pb.AddRoomResponse{Message: "Room added successfully"}, nil
}

func (h *HotelHandler) GetRoom(ctx context.Context, req *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
	room, err := h.service.GetRoom(req.RoomNumber)
	if err != nil {
		return nil, err
	}
	return &pb.GetRoomResponse{
		Room: &pb.Room{
			RoomNumber:   room.RoomNumber,
			Category:     room.Category,
			Availability: room.Availability,
			Price:        room.Price,
		},
	}, nil
}

func (h *HotelHandler) UpdateRoom(ctx context.Context, req *pb.UpdateRoomRequest) (*pb.UpdateRoomResponse, error) {
	room := domain.Room{
		RoomNumber:   req.Room.RoomNumber,
		Category:     req.Room.Category,
		Availability: req.Room.Availability,
		Price:        req.Room.Price,
	}
	err := h.service.UpdateRoom(room)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateRoomResponse{Message: "Room updated successfully"}, nil
}

func (h *HotelHandler) DeleteRoom(ctx context.Context, req *pb.DeleteRoomRequest) (*pb.DeleteRoomResponse, error) {
	err := h.service.DeleteRoom(req.RoomNumber)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteRoomResponse{Message: "Room deleted successfully"}, nil
}

func (h *HotelHandler) GetRooms(ctx context.Context, req *pb.GetRoomsRequest) (*pb.GetRoomsResponse, error) {
	rooms, err := h.service.GetRooms()
	if err != nil {
		return nil, err
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

	return &pb.GetRoomsResponse{Rooms: pbRooms}, nil
}

func (h *HotelHandler) CheckRoom(ctx context.Context, req *pb.CheckRoomRequest) (*pb.CheckRoomResponse, error) {
	available := h.service.RoomAvailable(req.RoomId)
	return &pb.CheckRoomResponse{Available: available}, nil
}
