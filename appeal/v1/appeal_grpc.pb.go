// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: appeal/v1/appeal.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Appeal_OpCreateAppeal_FullMethodName = "/api.appeal.v1.Appeal/OpCreateAppeal"
	Appeal_UpdateAppeal_FullMethodName   = "/api.appeal.v1.Appeal/UpdateAppeal"
	Appeal_DeleteAppeal_FullMethodName   = "/api.appeal.v1.Appeal/DeleteAppeal"
	Appeal_GetAppeal_FullMethodName      = "/api.appeal.v1.Appeal/GetAppeal"
	Appeal_ListAppeal_FullMethodName     = "/api.appeal.v1.Appeal/ListAppeal"
)

// AppealClient is the client API for Appeal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppealClient interface {
	OpCreateAppeal(ctx context.Context, in *OpCreateAppealRequest, opts ...grpc.CallOption) (*OpCreateAppealReply, error)
	UpdateAppeal(ctx context.Context, in *UpdateAppealRequest, opts ...grpc.CallOption) (*UpdateAppealReply, error)
	DeleteAppeal(ctx context.Context, in *DeleteAppealRequest, opts ...grpc.CallOption) (*DeleteAppealReply, error)
	GetAppeal(ctx context.Context, in *GetAppealRequest, opts ...grpc.CallOption) (*GetAppealReply, error)
	ListAppeal(ctx context.Context, in *ListAppealRequest, opts ...grpc.CallOption) (*ListAppealReply, error)
}

type appealClient struct {
	cc grpc.ClientConnInterface
}

func NewAppealClient(cc grpc.ClientConnInterface) AppealClient {
	return &appealClient{cc}
}

func (c *appealClient) OpCreateAppeal(ctx context.Context, in *OpCreateAppealRequest, opts ...grpc.CallOption) (*OpCreateAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OpCreateAppealReply)
	err := c.cc.Invoke(ctx, Appeal_OpCreateAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appealClient) UpdateAppeal(ctx context.Context, in *UpdateAppealRequest, opts ...grpc.CallOption) (*UpdateAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAppealReply)
	err := c.cc.Invoke(ctx, Appeal_UpdateAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appealClient) DeleteAppeal(ctx context.Context, in *DeleteAppealRequest, opts ...grpc.CallOption) (*DeleteAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAppealReply)
	err := c.cc.Invoke(ctx, Appeal_DeleteAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appealClient) GetAppeal(ctx context.Context, in *GetAppealRequest, opts ...grpc.CallOption) (*GetAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppealReply)
	err := c.cc.Invoke(ctx, Appeal_GetAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appealClient) ListAppeal(ctx context.Context, in *ListAppealRequest, opts ...grpc.CallOption) (*ListAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAppealReply)
	err := c.cc.Invoke(ctx, Appeal_ListAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppealServer is the server API for Appeal service.
// All implementations must embed UnimplementedAppealServer
// for forward compatibility.
type AppealServer interface {
	OpCreateAppeal(context.Context, *OpCreateAppealRequest) (*OpCreateAppealReply, error)
	UpdateAppeal(context.Context, *UpdateAppealRequest) (*UpdateAppealReply, error)
	DeleteAppeal(context.Context, *DeleteAppealRequest) (*DeleteAppealReply, error)
	GetAppeal(context.Context, *GetAppealRequest) (*GetAppealReply, error)
	ListAppeal(context.Context, *ListAppealRequest) (*ListAppealReply, error)
	mustEmbedUnimplementedAppealServer()
}

// UnimplementedAppealServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAppealServer struct{}

func (UnimplementedAppealServer) OpCreateAppeal(context.Context, *OpCreateAppealRequest) (*OpCreateAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpCreateAppeal not implemented")
}
func (UnimplementedAppealServer) UpdateAppeal(context.Context, *UpdateAppealRequest) (*UpdateAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppeal not implemented")
}
func (UnimplementedAppealServer) DeleteAppeal(context.Context, *DeleteAppealRequest) (*DeleteAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppeal not implemented")
}
func (UnimplementedAppealServer) GetAppeal(context.Context, *GetAppealRequest) (*GetAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppeal not implemented")
}
func (UnimplementedAppealServer) ListAppeal(context.Context, *ListAppealRequest) (*ListAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppeal not implemented")
}
func (UnimplementedAppealServer) mustEmbedUnimplementedAppealServer() {}
func (UnimplementedAppealServer) testEmbeddedByValue()                {}

// UnsafeAppealServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppealServer will
// result in compilation errors.
type UnsafeAppealServer interface {
	mustEmbedUnimplementedAppealServer()
}

func RegisterAppealServer(s grpc.ServiceRegistrar, srv AppealServer) {
	// If the following call pancis, it indicates UnimplementedAppealServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Appeal_ServiceDesc, srv)
}

func _Appeal_OpCreateAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpCreateAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppealServer).OpCreateAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Appeal_OpCreateAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppealServer).OpCreateAppeal(ctx, req.(*OpCreateAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appeal_UpdateAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppealServer).UpdateAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Appeal_UpdateAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppealServer).UpdateAppeal(ctx, req.(*UpdateAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appeal_DeleteAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppealServer).DeleteAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Appeal_DeleteAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppealServer).DeleteAppeal(ctx, req.(*DeleteAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appeal_GetAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppealServer).GetAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Appeal_GetAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppealServer).GetAppeal(ctx, req.(*GetAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appeal_ListAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppealServer).ListAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Appeal_ListAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppealServer).ListAppeal(ctx, req.(*ListAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Appeal_ServiceDesc is the grpc.ServiceDesc for Appeal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Appeal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.appeal.v1.Appeal",
	HandlerType: (*AppealServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpCreateAppeal",
			Handler:    _Appeal_OpCreateAppeal_Handler,
		},
		{
			MethodName: "UpdateAppeal",
			Handler:    _Appeal_UpdateAppeal_Handler,
		},
		{
			MethodName: "DeleteAppeal",
			Handler:    _Appeal_DeleteAppeal_Handler,
		},
		{
			MethodName: "GetAppeal",
			Handler:    _Appeal_GetAppeal_Handler,
		},
		{
			MethodName: "ListAppeal",
			Handler:    _Appeal_ListAppeal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appeal/v1/appeal.proto",
}
