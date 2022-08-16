// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: match/v1/match.proto

package matchv1

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

// MatchServiceClient is the client API for MatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchServiceClient interface {
	MatchPartnersWithRequest(ctx context.Context, in *MatchPartnersWithRequestInput, opts ...grpc.CallOption) (*MatchPartnersWithRequestOutput, error)
}

type matchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchServiceClient(cc grpc.ClientConnInterface) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) MatchPartnersWithRequest(ctx context.Context, in *MatchPartnersWithRequestInput, opts ...grpc.CallOption) (*MatchPartnersWithRequestOutput, error) {
	out := new(MatchPartnersWithRequestOutput)
	err := c.cc.Invoke(ctx, "/match.v1.matchService/MatchPartnersWithRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchServiceServer is the server API for MatchService service.
// All implementations should embed UnimplementedMatchServiceServer
// for forward compatibility
type MatchServiceServer interface {
	MatchPartnersWithRequest(context.Context, *MatchPartnersWithRequestInput) (*MatchPartnersWithRequestOutput, error)
}

// UnimplementedMatchServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMatchServiceServer struct {
}

func (UnimplementedMatchServiceServer) MatchPartnersWithRequest(context.Context, *MatchPartnersWithRequestInput) (*MatchPartnersWithRequestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MatchPartnersWithRequest not implemented")
}

// UnsafeMatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchServiceServer will
// result in compilation errors.
type UnsafeMatchServiceServer interface {
	mustEmbedUnimplementedMatchServiceServer()
}

func RegisterMatchServiceServer(s grpc.ServiceRegistrar, srv MatchServiceServer) {
	s.RegisterService(&MatchService_ServiceDesc, srv)
}

func _MatchService_MatchPartnersWithRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchPartnersWithRequestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).MatchPartnersWithRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/match.v1.matchService/MatchPartnersWithRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).MatchPartnersWithRequest(ctx, req.(*MatchPartnersWithRequestInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchService_ServiceDesc is the grpc.ServiceDesc for MatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "match.v1.matchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MatchPartnersWithRequest",
			Handler:    _MatchService_MatchPartnersWithRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "match/v1/match.proto",
}
