// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: app/grpc/proto/home-feeds/post_feed.proto

package proto

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
	PostService_GetPostFeeds_FullMethodName = "/feeds.PostService/GetPostFeeds"
	PostService_RefreshPost_FullMethodName  = "/feeds.PostService/RefreshPost"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	GetPostFeeds(ctx context.Context, in *GetPostFeedsRequest, opts ...grpc.CallOption) (*GetPostFeedsResponse, error)
	RefreshPost(ctx context.Context, in *RefreshPostRequest, opts ...grpc.CallOption) (*GetPostFeedsResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) GetPostFeeds(ctx context.Context, in *GetPostFeedsRequest, opts ...grpc.CallOption) (*GetPostFeedsResponse, error) {
	out := new(GetPostFeedsResponse)
	err := c.cc.Invoke(ctx, PostService_GetPostFeeds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) RefreshPost(ctx context.Context, in *RefreshPostRequest, opts ...grpc.CallOption) (*GetPostFeedsResponse, error) {
	out := new(GetPostFeedsResponse)
	err := c.cc.Invoke(ctx, PostService_RefreshPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	GetPostFeeds(context.Context, *GetPostFeedsRequest) (*GetPostFeedsResponse, error)
	RefreshPost(context.Context, *RefreshPostRequest) (*GetPostFeedsResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) GetPostFeeds(context.Context, *GetPostFeedsRequest) (*GetPostFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostFeeds not implemented")
}
func (UnimplementedPostServiceServer) RefreshPost(context.Context, *RefreshPostRequest) (*GetPostFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshPost not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_GetPostFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPostFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostFeeds(ctx, req.(*GetPostFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_RefreshPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).RefreshPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_RefreshPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).RefreshPost(ctx, req.(*RefreshPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feeds.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPostFeeds",
			Handler:    _PostService_GetPostFeeds_Handler,
		},
		{
			MethodName: "RefreshPost",
			Handler:    _PostService_RefreshPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/grpc/proto/home-feeds/post_feed.proto",
}
