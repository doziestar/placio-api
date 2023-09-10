package services

import (
	"context"
	"placio-pkg/grpc/proto"
)

type PostService interface {
	GetPostFeeds(ctx context.Context) (*proto.GetPostFeedsResponse, error)
	RefreshPostFeeds(ctx context.Context) (*proto.GetPostFeedsResponse, error)
	WatchPosts(ctx context.Context) (proto.PostService_WatchPostsClient, error)
}

type postService struct {
	service proto.PostServiceClient
}

func NewPostService(service proto.PostServiceClient) PostService {
	return &postService{service: service}
}

func (p *postService) GetPostFeeds(ctx context.Context) (*proto.GetPostFeedsResponse, error) {
	return p.service.GetPostFeeds(ctx, &proto.GetPostFeedsRequest{})
}

func (p *postService) RefreshPostFeeds(ctx context.Context) (*proto.GetPostFeedsResponse, error) {
	return p.service.RefreshPost(ctx, &proto.RefreshPostRequest{})
}

func (p *postService) WatchPosts(ctx context.Context) (proto.PostService_WatchPostsClient, error) {
	return p.service.WatchPosts(ctx)
}
