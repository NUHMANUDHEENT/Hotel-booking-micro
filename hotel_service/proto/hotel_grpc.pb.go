// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HotelServiceClient is the client API for HotelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelServiceClient interface {
	AddRoom(ctx context.Context, in *AddRoomRequest, opts ...grpc.CallOption) (*AddRoomResponse, error)
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error)
	UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomResponse, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error)
	GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error)
	CheckRoom(ctx context.Context, in *CheckRoomRequest, opts ...grpc.CallOption) (*CheckRoomResponse, error)
}

type hotelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelServiceClient(cc grpc.ClientConnInterface) HotelServiceClient {
	return &hotelServiceClient{cc}
}

func (c *hotelServiceClient) AddRoom(ctx context.Context, in *AddRoomRequest, opts ...grpc.CallOption) (*AddRoomResponse, error) {
	out := new(AddRoomResponse)
	err := c.cc.Invoke(ctx, "/hotel_service.HotelService/AddRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error) {
	out := new(GetRoomResponse)
	err := c.cc.Invoke(ctx, "/hotel_service.HotelService/GetRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomResponse, error) {
	out := new(UpdateRoomResponse)
	err := c.cc.Invoke(ctx, "/hotel_service.HotelService/UpdateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error) {
	out := new(DeleteRoomResponse)
	err := c.cc.Invoke(ctx, "/hotel_service.HotelService/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error) {
	out := new(GetRoomsResponse)
	err := c.cc.Invoke(ctx, "/hotel_service.HotelService/GetRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) CheckRoom(ctx context.Context, in *CheckRoomRequest, opts ...grpc.CallOption) (*CheckRoomResponse, error) {
	out := new(CheckRoomResponse)
	err := c.cc.Invoke(ctx, "/hotel_service.HotelService/CheckRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelServiceServer is the server API for HotelService service.
// All implementations must embed UnimplementedHotelServiceServer
// for forward compatibility
type HotelServiceServer interface {
	AddRoom(context.Context, *AddRoomRequest) (*AddRoomResponse, error)
	GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error)
	UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomResponse, error)
	DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error)
	GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error)
	CheckRoom(context.Context, *CheckRoomRequest) (*CheckRoomResponse, error)
	mustEmbedUnimplementedHotelServiceServer()
}

// UnimplementedHotelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHotelServiceServer struct {
}

func (UnimplementedHotelServiceServer) AddRoom(context.Context, *AddRoomRequest) (*AddRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoom not implemented")
}
func (UnimplementedHotelServiceServer) GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedHotelServiceServer) UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedHotelServiceServer) DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedHotelServiceServer) GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (UnimplementedHotelServiceServer) CheckRoom(context.Context, *CheckRoomRequest) (*CheckRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRoom not implemented")
}
func (UnimplementedHotelServiceServer) mustEmbedUnimplementedHotelServiceServer() {}

// UnsafeHotelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelServiceServer will
// result in compilation errors.
type UnsafeHotelServiceServer interface {
	mustEmbedUnimplementedHotelServiceServer()
}

func RegisterHotelServiceServer(s *grpc.Server, srv HotelServiceServer) {
	s.RegisterService(&_HotelService_serviceDesc, srv)
}

func _HotelService_AddRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).AddRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel_service.HotelService/AddRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).AddRoom(ctx, req.(*AddRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel_service.HotelService/GetRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel_service.HotelService/UpdateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).UpdateRoom(ctx, req.(*UpdateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel_service.HotelService/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel_service.HotelService/GetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetRooms(ctx, req.(*GetRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_CheckRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).CheckRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel_service.HotelService/CheckRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).CheckRoom(ctx, req.(*CheckRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hotel_service.HotelService",
	HandlerType: (*HotelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRoom",
			Handler:    _HotelService_AddRoom_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _HotelService_GetRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _HotelService_UpdateRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _HotelService_DeleteRoom_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _HotelService_GetRooms_Handler,
		},
		{
			MethodName: "CheckRoom",
			Handler:    _HotelService_CheckRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotel.proto",
}
