package cmd

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"placio-app/grpc/proto/home-feeds"

	"placio-app/domains/media"
	"placio-app/domains/posts"
	"placio-app/ent"
	"placio-app/utility"

	"github.com/asaskevich/EventBus"
	"github.com/cloudinary/cloudinary-go/v2"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedPostServiceServer
	postService posts.PostService
}

func (s *server) RefreshPost(ctx context.Context, req *proto.RefreshPostRequest) (*proto.GetPostFeedsResponse, error) {
	posts, err := s.postService.GetPostFeeds(ctx)
	if err != nil {
		return nil, err
	}

	// Convert feedPosts to protobuf's posts
	var pbPosts []*proto.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, &proto.Post{
			Id:        p.ID,
			Content:   p.Content,
			CreatedAt: p.CreatedAt.String(), // Convert time to string or google.protobuf.Timestamp
			UpdatedAt: p.UpdatedAt.String(),
			Privacy: func() proto.Post_PrivacyType {
				switch p.Privacy {
				case "PUBLIC":
					return *proto.Post_PRIVATE.Enum()
				case "PRIVATE":
					return *proto.Post_PUBLIC.Enum()
				default:
					return *proto.Post_UNKNOWN.Enum()
				}
			}(),
			// Edges: &pb.Post_Edge{
			// 	Comments: func () []*pb.Comment {
			// },
		},
		)
	}

	return &proto.GetPostFeedsResponse{Posts: pbPosts}, nil
}

func (s *server) GetPostFeeds(ctx context.Context, req *proto.GetPostFeedsRequest) (*proto.GetPostFeedsResponse, error) {
	feedPosts, err := s.postService.GetPostFeeds(ctx)
	if err != nil {
		return nil, err
	}

	// Convert feedPosts to protobuf's posts
	var pbPosts []*proto.Post
	for _, p := range feedPosts {
		pbPosts = append(pbPosts, &proto.Post{
			Id:        p.ID,
			Content:   p.Content,
			CreatedAt: p.CreatedAt.String(), // Convert time to string or google.protobuf.Timestamp
			UpdatedAt: p.UpdatedAt.String(),
			Privacy: func() proto.Post_PrivacyType {
				switch p.Privacy {
				case "PUBLIC":
					return *proto.Post_PRIVATE.Enum()
				case "PRIVATE":
					return *proto.Post_PUBLIC.Enum()
				default:
					return *proto.Post_UNKNOWN.Enum()
				}
			}(),
			// Edges: &pb.Post_Edge{
			// 	Comments: func () []*pb.Comment {
			// },
		},
		)
	}

	return &proto.GetPostFeedsResponse{Posts: pbPosts}, nil
}

func ServeGRPC(client *ent.Client) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	redisClient := utility.NewRedisClient("redis://default:a3677c1a7b84402eb34efd55ad3cf059@golden-colt-33790.upstash.io:33790", 0)
	_ = redisClient.ConnectRedis()

	cld, _ := cloudinary.NewFromParams("placio", "312498583624125", "k4XSQwWuhi3Vy7QAw7Qn0mUaW0s")
	eventBus := EventBus.New()

	mediaService := media.NewMediaService(client, cld)

	s := grpc.NewServer()
	postSvc := posts.NewPostService(client, redisClient, mediaService, eventBus)
	proto.RegisterPostServiceServer(s, &server{postService: postSvc})

	log.Println("gRPC server started on :50051")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	s.GracefulStop()
	log.Println("gRPC server stopped gracefully")
}
