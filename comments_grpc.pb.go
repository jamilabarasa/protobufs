// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/comments.proto

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
	Commentservice_GetComment_FullMethodName = "/proto.commentservice/GetComment"
)

// CommentserviceClient is the client API for Commentservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The comments gRPC service
type CommentserviceClient interface {
	GetComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
}

type commentserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentserviceClient(cc grpc.ClientConnInterface) CommentserviceClient {
	return &commentserviceClient{cc}
}

func (c *commentserviceClient) GetComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, Commentservice_GetComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentserviceServer is the server API for Commentservice service.
// All implementations must embed UnimplementedCommentserviceServer
// for forward compatibility.
//
// The comments gRPC service
type CommentserviceServer interface {
	GetComment(context.Context, *CommentRequest) (*CommentResponse, error)
	mustEmbedUnimplementedCommentserviceServer()
}

// UnimplementedCommentserviceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentserviceServer struct{}

func (UnimplementedCommentserviceServer) GetComment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCommentserviceServer) mustEmbedUnimplementedCommentserviceServer() {}
func (UnimplementedCommentserviceServer) testEmbeddedByValue()                        {}

// UnsafeCommentserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentserviceServer will
// result in compilation errors.
type UnsafeCommentserviceServer interface {
	mustEmbedUnimplementedCommentserviceServer()
}

func RegisterCommentserviceServer(s grpc.ServiceRegistrar, srv CommentserviceServer) {
	// If the following call pancis, it indicates UnimplementedCommentserviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Commentservice_ServiceDesc, srv)
}

func _Commentservice_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentserviceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Commentservice_GetComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentserviceServer).GetComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Commentservice_ServiceDesc is the grpc.ServiceDesc for Commentservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Commentservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.commentservice",
	HandlerType: (*CommentserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComment",
			Handler:    _Commentservice_GetComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/comments.proto",
}
