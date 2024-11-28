// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/reactions.proto

package __

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
	Reactionservice_GetReaction_FullMethodName = "/proto.reactionservice/GetReaction"
)

// ReactionserviceClient is the client API for Reactionservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The reactions gRPC service
type ReactionserviceClient interface {
	GetReaction(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error)
}

type reactionserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewReactionserviceClient(cc grpc.ClientConnInterface) ReactionserviceClient {
	return &reactionserviceClient{cc}
}

func (c *reactionserviceClient) GetReaction(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReactionResponse)
	err := c.cc.Invoke(ctx, Reactionservice_GetReaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReactionserviceServer is the server API for Reactionservice service.
// All implementations must embed UnimplementedReactionserviceServer
// for forward compatibility.
//
// The reactions gRPC service
type ReactionserviceServer interface {
	GetReaction(context.Context, *ReactionRequest) (*ReactionResponse, error)
	mustEmbedUnimplementedReactionserviceServer()
}

// UnimplementedReactionserviceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReactionserviceServer struct{}

func (UnimplementedReactionserviceServer) GetReaction(context.Context, *ReactionRequest) (*ReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReaction not implemented")
}
func (UnimplementedReactionserviceServer) mustEmbedUnimplementedReactionserviceServer() {}
func (UnimplementedReactionserviceServer) testEmbeddedByValue()                         {}

// UnsafeReactionserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReactionserviceServer will
// result in compilation errors.
type UnsafeReactionserviceServer interface {
	mustEmbedUnimplementedReactionserviceServer()
}

func RegisterReactionserviceServer(s grpc.ServiceRegistrar, srv ReactionserviceServer) {
	// If the following call pancis, it indicates UnimplementedReactionserviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Reactionservice_ServiceDesc, srv)
}

func _Reactionservice_GetReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionserviceServer).GetReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reactionservice_GetReaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionserviceServer).GetReaction(ctx, req.(*ReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reactionservice_ServiceDesc is the grpc.ServiceDesc for Reactionservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reactionservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.reactionservice",
	HandlerType: (*ReactionserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReaction",
			Handler:    _Reactionservice_GetReaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reactions.proto",
}
