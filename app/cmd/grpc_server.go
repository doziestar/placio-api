package cmd

import (
	"context"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"os"
	"os/signal"
	"placio-api/grpc/proto"
	"time"

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
	eventBus    EventBus.Bus
	postService posts.PostService
}

func NewServer(postService posts.PostService, eventBus EventBus.Bus) proto.PostServiceServer {
	return &server{postService: postService, eventBus: eventBus}
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

func (s *server) WatchPosts(stream proto.PostService_WatchPostsServer) error {
	postsUpdated := make(chan bool, 100)

	// Subscription to listen for new posts
	log.Println("Client connected to WatchPosts stream.")
	err := s.eventBus.Subscribe("post:created", func(post *ent.Post) {
		log.Println("New post created. Broadcasting to clients...")
		postsUpdated <- true
	})
	log.Println("Subscribed to post:created event.")

	if err != nil {
		log.Fatalf("failed to subscribe to event: %v", err)
	}

	defer s.eventBus.Unsubscribe("post:created", func(post *ent.Post) {
		log.Println("Unsubscribed from post:created event.")
	})

	for {
		select {
		case <-postsUpdated:
			log.Println("Sending new posts to client...")
			posts, err := s.postService.GetPostFeeds(stream.Context())
			log.Println("posts: ", posts)
			if err != nil {
				return err
			}

			var pbPosts []*proto.Post
			for _, p := range posts {
				pbPosts = append(pbPosts, &proto.Post{
					Id:        p.ID,
					Content:   p.Content,
					CreatedAt: p.CreatedAt.String(),
					UpdatedAt: p.UpdatedAt.String(),
					// Other fields...
				})
			}

			if err := stream.Send(&proto.GetPostFeedsResponse{Posts: pbPosts}); err != nil {
				return err
			}

		case <-stream.Context().Done():
			log.Println("Client disconnected from WatchPosts stream.")
			return stream.Context().Err()
		}
	}
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

	ka := keepalive.ServerParameters{
		Time:    10 * time.Hour,
		Timeout: 5 * time.Hour,
	}
	s := grpc.NewServer(grpc.KeepaliveParams(ka))

	postSvc := posts.NewPostService(client, redisClient, mediaService, eventBus)
	proto.RegisterPostServiceServer(s, &server{postService: postSvc, eventBus: eventBus})

	log.Println("gRPC server started on :50051")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	err = eventBus.Subscribe("post:created", func(post *ent.Post) {
		log.Println("Post created event triggered!")
	})

	if err != nil {
		log.Fatalf("failed to subscribe to event: %v", err)
	}

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	s.GracefulStop()
	log.Println("gRPC server stopped gracefully")
}
