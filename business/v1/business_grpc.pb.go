// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: business/v1/business.proto

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
	Business_GreateReply_FullMethodName  = "/api.business.v1.Business/GreateReply"
	Business_GreateAppeal_FullMethodName = "/api.business.v1.Business/GreateAppeal"
)

// BusinessClient is the client API for Business service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusinessClient interface {
	GreateReply(ctx context.Context, in *CreateReplyReq, opts ...grpc.CallOption) (*CreateReplyResp, error)
	GreateAppeal(ctx context.Context, in *CreateAppealRequest, opts ...grpc.CallOption) (*CreateAppealReply, error)
}

type businessClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessClient(cc grpc.ClientConnInterface) BusinessClient {
	return &businessClient{cc}
}

func (c *businessClient) GreateReply(ctx context.Context, in *CreateReplyReq, opts ...grpc.CallOption) (*CreateReplyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateReplyResp)
	err := c.cc.Invoke(ctx, Business_GreateReply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessClient) GreateAppeal(ctx context.Context, in *CreateAppealRequest, opts ...grpc.CallOption) (*CreateAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAppealReply)
	err := c.cc.Invoke(ctx, Business_GreateAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessServer is the server API for Business service.
// All implementations must embed UnimplementedBusinessServer
// for forward compatibility.
type BusinessServer interface {
	GreateReply(context.Context, *CreateReplyReq) (*CreateReplyResp, error)
	GreateAppeal(context.Context, *CreateAppealRequest) (*CreateAppealReply, error)
	mustEmbedUnimplementedBusinessServer()
}

// UnimplementedBusinessServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBusinessServer struct{}

func (UnimplementedBusinessServer) GreateReply(context.Context, *CreateReplyReq) (*CreateReplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreateReply not implemented")
}
func (UnimplementedBusinessServer) GreateAppeal(context.Context, *CreateAppealRequest) (*CreateAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreateAppeal not implemented")
}
func (UnimplementedBusinessServer) mustEmbedUnimplementedBusinessServer() {}
func (UnimplementedBusinessServer) testEmbeddedByValue()                  {}

// UnsafeBusinessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessServer will
// result in compilation errors.
type UnsafeBusinessServer interface {
	mustEmbedUnimplementedBusinessServer()
}

func RegisterBusinessServer(s grpc.ServiceRegistrar, srv BusinessServer) {
	// If the following call pancis, it indicates UnimplementedBusinessServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Business_ServiceDesc, srv)
}

func _Business_GreateReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServer).GreateReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_GreateReply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServer).GreateReply(ctx, req.(*CreateReplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Business_GreateAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServer).GreateAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_GreateAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServer).GreateAppeal(ctx, req.(*CreateAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Business_ServiceDesc is the grpc.ServiceDesc for Business service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Business_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.business.v1.Business",
	HandlerType: (*BusinessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreateReply",
			Handler:    _Business_GreateReply_Handler,
		},
		{
			MethodName: "GreateAppeal",
			Handler:    _Business_GreateAppeal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "business/v1/business.proto",
}
