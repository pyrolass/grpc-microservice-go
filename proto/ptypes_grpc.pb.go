// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/ptypes.proto

package types

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
	DriverMessageService_InsertDriverLog_FullMethodName = "/DriverMessageService/InsertDriverLog"
)

// DriverMessageServiceClient is the client API for DriverMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverMessageServiceClient interface {
	InsertDriverLog(ctx context.Context, in *InsertDriverLogRequest, opts ...grpc.CallOption) (*InsertDriverLogResponse, error)
}

type driverMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverMessageServiceClient(cc grpc.ClientConnInterface) DriverMessageServiceClient {
	return &driverMessageServiceClient{cc}
}

func (c *driverMessageServiceClient) InsertDriverLog(ctx context.Context, in *InsertDriverLogRequest, opts ...grpc.CallOption) (*InsertDriverLogResponse, error) {
	out := new(InsertDriverLogResponse)
	err := c.cc.Invoke(ctx, DriverMessageService_InsertDriverLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverMessageServiceServer is the server API for DriverMessageService service.
// All implementations must embed UnimplementedDriverMessageServiceServer
// for forward compatibility
type DriverMessageServiceServer interface {
	InsertDriverLog(context.Context, *InsertDriverLogRequest) (*InsertDriverLogResponse, error)
	mustEmbedUnimplementedDriverMessageServiceServer()
}

// UnimplementedDriverMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDriverMessageServiceServer struct {
}

func (UnimplementedDriverMessageServiceServer) InsertDriverLog(context.Context, *InsertDriverLogRequest) (*InsertDriverLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertDriverLog not implemented")
}
func (UnimplementedDriverMessageServiceServer) mustEmbedUnimplementedDriverMessageServiceServer() {}

// UnsafeDriverMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverMessageServiceServer will
// result in compilation errors.
type UnsafeDriverMessageServiceServer interface {
	mustEmbedUnimplementedDriverMessageServiceServer()
}

func RegisterDriverMessageServiceServer(s grpc.ServiceRegistrar, srv DriverMessageServiceServer) {
	s.RegisterService(&DriverMessageService_ServiceDesc, srv)
}

func _DriverMessageService_InsertDriverLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertDriverLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverMessageServiceServer).InsertDriverLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverMessageService_InsertDriverLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverMessageServiceServer).InsertDriverLog(ctx, req.(*InsertDriverLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverMessageService_ServiceDesc is the grpc.ServiceDesc for DriverMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DriverMessageService",
	HandlerType: (*DriverMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertDriverLog",
			Handler:    _DriverMessageService_InsertDriverLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ptypes.proto",
}
