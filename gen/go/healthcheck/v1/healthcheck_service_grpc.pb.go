// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: botapi/healthcheck/v1/healthcheck_service.proto

package healthcheck

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthcheckServiceClient is the client API for HealthcheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthcheckServiceClient interface {
	// check method returns any string to validate that service is running and
	// grpc interface operates
	Check(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CheckResponse, error)
}

type healthcheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthcheckServiceClient(cc grpc.ClientConnInterface) HealthcheckServiceClient {
	return &healthcheckServiceClient{cc}
}

func (c *healthcheckServiceClient) Check(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/botapi.healthcheck.v1.HealthcheckService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthcheckServiceServer is the server API for HealthcheckService service.
// All implementations must embed UnimplementedHealthcheckServiceServer
// for forward compatibility
type HealthcheckServiceServer interface {
	// check method returns any string to validate that service is running and
	// grpc interface operates
	Check(context.Context, *emptypb.Empty) (*CheckResponse, error)
	mustEmbedUnimplementedHealthcheckServiceServer()
}

// UnimplementedHealthcheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthcheckServiceServer struct {
}

func (UnimplementedHealthcheckServiceServer) Check(context.Context, *emptypb.Empty) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedHealthcheckServiceServer) mustEmbedUnimplementedHealthcheckServiceServer() {}

// UnsafeHealthcheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthcheckServiceServer will
// result in compilation errors.
type UnsafeHealthcheckServiceServer interface {
	mustEmbedUnimplementedHealthcheckServiceServer()
}

func RegisterHealthcheckServiceServer(s grpc.ServiceRegistrar, srv HealthcheckServiceServer) {
	s.RegisterService(&HealthcheckService_ServiceDesc, srv)
}

func _HealthcheckService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthcheckServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botapi.healthcheck.v1.HealthcheckService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthcheckServiceServer).Check(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthcheckService_ServiceDesc is the grpc.ServiceDesc for HealthcheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthcheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "botapi.healthcheck.v1.HealthcheckService",
	HandlerType: (*HealthcheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _HealthcheckService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "botapi/healthcheck/v1/healthcheck_service.proto",
}
