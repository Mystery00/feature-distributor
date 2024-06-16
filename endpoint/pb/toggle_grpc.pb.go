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
	ToggleService_ListToggle_FullMethodName      = "/ToggleService/ListToggle"
	ToggleService_GetToggle_FullMethodName       = "/ToggleService/GetToggle"
	ToggleService_GetBoolToggle_FullMethodName   = "/ToggleService/GetBoolToggle"
	ToggleService_GetStringToggle_FullMethodName = "/ToggleService/GetStringToggle"
	ToggleService_GetFloatToggle_FullMethodName  = "/ToggleService/GetFloatToggle"
	ToggleService_GetIntToggle_FullMethodName    = "/ToggleService/GetIntToggle"
	ToggleService_GetJsonToggle_FullMethodName   = "/ToggleService/GetJsonToggle"
)

// ToggleServiceClient is the client API for ToggleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToggleServiceClient interface {
	ListToggle(ctx context.Context, in *ListToggleRequest, opts ...grpc.CallOption) (*ListToggleResponse, error)
	GetToggle(ctx context.Context, in *GetToggleRequest, opts ...grpc.CallOption) (*Toggle, error)
	GetBoolToggle(ctx context.Context, in *BoolToggleRequest, opts ...grpc.CallOption) (*BoolToggleResponse, error)
	GetStringToggle(ctx context.Context, in *StringToggleRequest, opts ...grpc.CallOption) (*StringToggleResponse, error)
	GetFloatToggle(ctx context.Context, in *FloatToggleRequest, opts ...grpc.CallOption) (*FloatToggleResponse, error)
	GetIntToggle(ctx context.Context, in *IntToggleRequest, opts ...grpc.CallOption) (*IntToggleResponse, error)
	GetJsonToggle(ctx context.Context, in *JsonToggleRequest, opts ...grpc.CallOption) (*JsonToggleResponse, error)
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

func (c *toggleServiceClient) GetBoolToggle(ctx context.Context, in *BoolToggleRequest, opts ...grpc.CallOption) (*BoolToggleResponse, error) {
	out := new(BoolToggleResponse)
	err := c.cc.Invoke(ctx, ToggleService_GetBoolToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toggleServiceClient) GetStringToggle(ctx context.Context, in *StringToggleRequest, opts ...grpc.CallOption) (*StringToggleResponse, error) {
	out := new(StringToggleResponse)
	err := c.cc.Invoke(ctx, ToggleService_GetStringToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toggleServiceClient) GetFloatToggle(ctx context.Context, in *FloatToggleRequest, opts ...grpc.CallOption) (*FloatToggleResponse, error) {
	out := new(FloatToggleResponse)
	err := c.cc.Invoke(ctx, ToggleService_GetFloatToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toggleServiceClient) GetIntToggle(ctx context.Context, in *IntToggleRequest, opts ...grpc.CallOption) (*IntToggleResponse, error) {
	out := new(IntToggleResponse)
	err := c.cc.Invoke(ctx, ToggleService_GetIntToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toggleServiceClient) GetJsonToggle(ctx context.Context, in *JsonToggleRequest, opts ...grpc.CallOption) (*JsonToggleResponse, error) {
	out := new(JsonToggleResponse)
	err := c.cc.Invoke(ctx, ToggleService_GetJsonToggle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToggleServiceServer is the toggle API for ToggleService service.
// All implementations must embed UnimplementedToggleServiceServer
// for forward compatibility
type ToggleServiceServer interface {
	ListToggle(context.Context, *ListToggleRequest) (*ListToggleResponse, error)
	GetToggle(context.Context, *GetToggleRequest) (*Toggle, error)
	GetBoolToggle(context.Context, *BoolToggleRequest) (*BoolToggleResponse, error)
	GetStringToggle(context.Context, *StringToggleRequest) (*StringToggleResponse, error)
	GetFloatToggle(context.Context, *FloatToggleRequest) (*FloatToggleResponse, error)
	GetIntToggle(context.Context, *IntToggleRequest) (*IntToggleResponse, error)
	GetJsonToggle(context.Context, *JsonToggleRequest) (*JsonToggleResponse, error)
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
func (UnimplementedToggleServiceServer) GetBoolToggle(context.Context, *BoolToggleRequest) (*BoolToggleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoolToggle not implemented")
}
func (UnimplementedToggleServiceServer) GetStringToggle(context.Context, *StringToggleRequest) (*StringToggleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStringToggle not implemented")
}
func (UnimplementedToggleServiceServer) GetFloatToggle(context.Context, *FloatToggleRequest) (*FloatToggleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFloatToggle not implemented")
}
func (UnimplementedToggleServiceServer) GetIntToggle(context.Context, *IntToggleRequest) (*IntToggleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntToggle not implemented")
}
func (UnimplementedToggleServiceServer) GetJsonToggle(context.Context, *JsonToggleRequest) (*JsonToggleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJsonToggle not implemented")
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

func _ToggleService_GetBoolToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoolToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).GetBoolToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_GetBoolToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).GetBoolToggle(ctx, req.(*BoolToggleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToggleService_GetStringToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).GetStringToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_GetStringToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).GetStringToggle(ctx, req.(*StringToggleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToggleService_GetFloatToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FloatToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).GetFloatToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_GetFloatToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).GetFloatToggle(ctx, req.(*FloatToggleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToggleService_GetIntToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).GetIntToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_GetIntToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).GetIntToggle(ctx, req.(*IntToggleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToggleService_GetJsonToggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsonToggleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToggleServiceServer).GetJsonToggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToggleService_GetJsonToggle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToggleServiceServer).GetJsonToggle(ctx, req.(*JsonToggleRequest))
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
			MethodName: "GetBoolToggle",
			Handler:    _ToggleService_GetBoolToggle_Handler,
		},
		{
			MethodName: "GetStringToggle",
			Handler:    _ToggleService_GetStringToggle_Handler,
		},
		{
			MethodName: "GetFloatToggle",
			Handler:    _ToggleService_GetFloatToggle_Handler,
		},
		{
			MethodName: "GetIntToggle",
			Handler:    _ToggleService_GetIntToggle_Handler,
		},
		{
			MethodName: "GetJsonToggle",
			Handler:    _ToggleService_GetJsonToggle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/protobuf/toggle.proto",
}
