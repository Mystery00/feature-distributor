// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: common/protobuf/toggle.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ToggleService_ListToggle_FullMethodName     = "/ToggleService/ListToggle"
	ToggleService_GetToggle_FullMethodName      = "/ToggleService/GetToggle"
	ToggleService_GetToggleValue_FullMethodName = "/ToggleService/GetToggleValue"
	ToggleService_CreateToggle_FullMethodName   = "/ToggleService/CreateToggle"
	ToggleService_DeleteToggle_FullMethodName   = "/ToggleService/DeleteToggle"
)

// ToggleServiceClient is the client API for ToggleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToggleServiceClient interface {
	ListToggle(ctx context.Context, in *ListToggleRequest, opts ...grpc.CallOption) (*ListToggleResponse, error)
	GetToggle(ctx context.Context, in *GetToggleRequest, opts ...grpc.CallOption) (*Toggle, error)
	GetToggleValue(ctx context.Context, in *GetToggleValueRequest, opts ...grpc.CallOption) (*GetToggleValueResponse, error)
	CreateToggle(ctx context.Context, in *CreateToggleRequest, opts ...grpc.CallOption) (*ToggleOperationResponse, error)
	DeleteToggle(ctx context.Context, in *GetToggleRequest, opts ...grpc.CallOption) (*ToggleOperationResponse, error)
}

type toggleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToggleServiceClient(cc grpc.ClientConnInterface) ToggleServiceClient {
	return &toggleServiceClient{cc}
}

func (c *toggleServiceClient) ListToggle(ctx context.Context, in *ListToggleRequest, opts ...grpc.CallOption) (*ListToggleResponse, error) {
	out := new(ListToggleResponse)
	err := c.cc.Invoke(ctx, ToggleService_ListToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toggleServiceClient) GetToggle(ctx context.Context, in *GetToggleRequest, opts ...grpc.CallOption) (*Toggle, error) {
	out := new(Toggle)
	err := c.cc.Invoke(ctx, ToggleService_GetToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toggleServiceClient) GetToggleValue(ctx context.Context, in *GetToggleValueRequest, opts ...grpc.CallOption) (*GetToggleValueResponse, error) {
	out := new(GetToggleValueResponse)
	err := c.cc.Invoke(ctx, ToggleService_GetToggleValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toggleServiceClient) CreateToggle(ctx context.Context, in *CreateToggleRequest, opts ...grpc.CallOption) (*ToggleOperationResponse, error) {
	out := new(ToggleOperationResponse)
	err := c.cc.Invoke(ctx, ToggleService_CreateToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toggleServiceClient) DeleteToggle(ctx context.Context, in *GetToggleRequest, opts ...grpc.CallOption) (*ToggleOperationResponse, error) {
	out := new(ToggleOperationResponse)
	err := c.cc.Invoke(ctx, ToggleService_DeleteToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToggleServiceServer is the server API for ToggleService service.
// All implementations must embed UnimplementedToggleServiceServer
// for forward compatibility
type ToggleServiceServer interface {
	ListToggle(context.Context, *ListToggleRequest) (*ListToggleResponse, error)
	GetToggle(context.Context, *GetToggleRequest) (*Toggle, error)
	GetToggleValue(context.Context, *GetToggleValueRequest) (*GetToggleValueResponse, error)
	CreateToggle(context.Context, *CreateToggleRequest) (*ToggleOperationResponse, error)
	DeleteToggle(context.Context, *GetToggleRequest) (*ToggleOperationResponse, error)
	mustEmbedUnimplementedToggleServiceServer()
}

// UnimplementedToggleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedToggleServiceServer struct {
}

func (UnimplementedToggleServiceServer) ListToggle(context.Context, *ListToggleRequest) (*ListToggleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListToggle not implemented")
}
func (UnimplementedToggleServiceServer) GetToggle(context.Context, *GetToggleRequest) (*Toggle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToggle not implemented")
}
func (UnimplementedToggleServiceServer) GetToggleValue(context.Context, *GetToggleValueRequest) (*GetToggleValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToggleValue not implemented")
}
func (UnimplementedToggleServiceServer) CreateToggle(context.Context, *CreateToggleRequest) (*ToggleOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToggle not implemented")
}
func (UnimplementedToggleServiceServer) DeleteToggle(context.Context, *GetToggleRequest) (*ToggleOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToggle not implemented")
}
func (UnimplementedToggleServiceServer) mustEmbedUnimplementedToggleServiceServer() {}

// UnsafeToggleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToggleServiceServer will
// result in compilation errors.
type UnsafeToggleServiceServer interface {
	mustEmbedUnimplementedToggleServiceServer()
}

func RegisterToggleServiceServer(s grpc.ServiceRegistrar, srv ToggleServiceServer) {
	s.RegisterService(&ToggleService_ServiceDesc, srv)
}

func _ToggleService_ListToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).ListToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_ListToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).ListToggle(ctx, req.(*ListToggleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToggleService_GetToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).GetToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_GetToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).GetToggle(ctx, req.(*GetToggleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToggleService_GetToggleValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetToggleValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).GetToggleValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_GetToggleValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).GetToggleValue(ctx, req.(*GetToggleValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToggleService_CreateToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).CreateToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_CreateToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).CreateToggle(ctx, req.(*CreateToggleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToggleService_DeleteToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).DeleteToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_DeleteToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).DeleteToggle(ctx, req.(*GetToggleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ToggleService_ServiceDesc is the grpc.ServiceDesc for ToggleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToggleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ToggleService",
	HandlerType: (*ToggleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListToggle",
			Handler:    _ToggleService_ListToggle_Handler,
		},
		{
			MethodName: "GetToggle",
			Handler:    _ToggleService_GetToggle_Handler,
		},
		{
			MethodName: "GetToggleValue",
			Handler:    _ToggleService_GetToggleValue_Handler,
		},
		{
			MethodName: "CreateToggle",
			Handler:    _ToggleService_CreateToggle_Handler,
		},
		{
			MethodName: "DeleteToggle",
			Handler:    _ToggleService_DeleteToggle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/protobuf/toggle.proto",
}
