// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.1
// source: food-svc.proto

package proto_gen

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FoodClient is the client API for Food service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoodClient interface {
	GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error)
	CreateFood(ctx context.Context, in *FoodDto, opts ...grpc.CallOption) (*FoodDto, error)
	GetFoodById(ctx context.Context, in *PayloadIdString, opts ...grpc.CallOption) (*FoodDto, error)
	ListFood(ctx context.Context, in *ListOptions, opts ...grpc.CallOption) (*ListFoodDto, error)
	UpdateFood(ctx context.Context, in *FoodDto, opts ...grpc.CallOption) (*FoodDto, error)
	DeleteFood(ctx context.Context, in *UpdatePayload, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
}

type foodClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodClient(cc grpc.ClientConnInterface) FoodClient {
	return &foodClient{cc}
}

func (c *foodClient) GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := c.cc.Invoke(ctx, "/foodSvc.Food/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) CreateFood(ctx context.Context, in *FoodDto, opts ...grpc.CallOption) (*FoodDto, error) {
	out := new(FoodDto)
	err := c.cc.Invoke(ctx, "/foodSvc.Food/CreateFood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) GetFoodById(ctx context.Context, in *PayloadIdString, opts ...grpc.CallOption) (*FoodDto, error) {
	out := new(FoodDto)
	err := c.cc.Invoke(ctx, "/foodSvc.Food/GetFoodById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) ListFood(ctx context.Context, in *ListOptions, opts ...grpc.CallOption) (*ListFoodDto, error) {
	out := new(ListFoodDto)
	err := c.cc.Invoke(ctx, "/foodSvc.Food/ListFood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) UpdateFood(ctx context.Context, in *FoodDto, opts ...grpc.CallOption) (*FoodDto, error) {
	out := new(FoodDto)
	err := c.cc.Invoke(ctx, "/foodSvc.Food/UpdateFood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) DeleteFood(ctx context.Context, in *UpdatePayload, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/foodSvc.Food/DeleteFood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoodServer is the server API for Food service.
// All implementations must embed UnimplementedFoodServer
// for forward compatibility
type FoodServer interface {
	GetHealth(context.Context, *empty.Empty) (*Health, error)
	CreateFood(context.Context, *FoodDto) (*FoodDto, error)
	GetFoodById(context.Context, *PayloadIdString) (*FoodDto, error)
	ListFood(context.Context, *ListOptions) (*ListFoodDto, error)
	UpdateFood(context.Context, *FoodDto) (*FoodDto, error)
	DeleteFood(context.Context, *UpdatePayload) (*wrappers.BoolValue, error)
	mustEmbedUnimplementedFoodServer()
}

// UnimplementedFoodServer must be embedded to have forward compatible implementations.
type UnimplementedFoodServer struct {
}

func (UnimplementedFoodServer) GetHealth(context.Context, *empty.Empty) (*Health, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedFoodServer) CreateFood(context.Context, *FoodDto) (*FoodDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFood not implemented")
}
func (UnimplementedFoodServer) GetFoodById(context.Context, *PayloadIdString) (*FoodDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoodById not implemented")
}
func (UnimplementedFoodServer) ListFood(context.Context, *ListOptions) (*ListFoodDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFood not implemented")
}
func (UnimplementedFoodServer) UpdateFood(context.Context, *FoodDto) (*FoodDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFood not implemented")
}
func (UnimplementedFoodServer) DeleteFood(context.Context, *UpdatePayload) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFood not implemented")
}
func (UnimplementedFoodServer) mustEmbedUnimplementedFoodServer() {}

// UnsafeFoodServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoodServer will
// result in compilation errors.
type UnsafeFoodServer interface {
	mustEmbedUnimplementedFoodServer()
}

func RegisterFoodServer(s grpc.ServiceRegistrar, srv FoodServer) {
	s.RegisterService(&Food_ServiceDesc, srv)
}

func _Food_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foodSvc.Food/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).GetHealth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_CreateFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FoodDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).CreateFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foodSvc.Food/CreateFood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).CreateFood(ctx, req.(*FoodDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_GetFoodById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadIdString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).GetFoodById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foodSvc.Food/GetFoodById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).GetFoodById(ctx, req.(*PayloadIdString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_ListFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).ListFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foodSvc.Food/ListFood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).ListFood(ctx, req.(*ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_UpdateFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FoodDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).UpdateFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foodSvc.Food/UpdateFood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).UpdateFood(ctx, req.(*FoodDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_DeleteFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).DeleteFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foodSvc.Food/DeleteFood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).DeleteFood(ctx, req.(*UpdatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

// Food_ServiceDesc is the grpc.ServiceDesc for Food service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Food_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "foodSvc.Food",
	HandlerType: (*FoodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _Food_GetHealth_Handler,
		},
		{
			MethodName: "CreateFood",
			Handler:    _Food_CreateFood_Handler,
		},
		{
			MethodName: "GetFoodById",
			Handler:    _Food_GetFoodById_Handler,
		},
		{
			MethodName: "ListFood",
			Handler:    _Food_ListFood_Handler,
		},
		{
			MethodName: "UpdateFood",
			Handler:    _Food_UpdateFood_Handler,
		},
		{
			MethodName: "DeleteFood",
			Handler:    _Food_DeleteFood_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "food-svc.proto",
}
