package cmd

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"placio-pkg/grpc/proto"
	"placio-pkg/kafka"
	"time"

	"google.golang.org/grpc/keepalive"

	"placio-app/domains/media"
	"placio-app/domains/posts"
	"placio-app/ent"
	"placio-app/utility"

	"github.com/cloudinary/cloudinary-go/v2"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedPostServiceServer
	producer    *kafka.Producer
	consumer    *kafka.KafkaConsumer
	postService posts.PostService
}

func NewServer(postService posts.PostService, producer *kafka.Producer, consumer *kafka.KafkaConsumer) proto.PostServiceServer {
	return &server{postService: postService, producer: producer, consumer: consumer}
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
			Edges: &proto.Post_Edge{
				User: &proto.Post_User{
					Id: 	  p.Edges.User.ID,
					Username: p.Edges.User.Username,
					Name: 	  p.Edges.User.Name,
					Picture: p.Edges.User.Picture,
				},
				Comments: func() []*proto.Post_Comment {
					var pbComments []*proto.Post_Comment
					for _, c := range p.Edges.Comments {
						pbComments = append(pbComments, &proto.Post_Comment{
							Id:        c.ID,
							Content:   c.Content,
							CreatedAt: c.CreatedAt.String(),
							UpdatedAt: c.UpdatedAt.String(),
							Edges: &proto.Post_CommentEdge{
								User: &proto.Post_User{
									Id: 	  c.Edges.User.ID,
									Username: c.Edges.User.Username,
									Name: 	  c.Edges.User.Name,
									Picture:  c.Edges.User.Picture,
								},
							},
						})
					}

					return pbComments
				}(),
			},
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

	log.Println("Client connected to WatchPosts stream.")

	// Start the Kafka consumer
	//go s.consumer.Start(postsUpdated)

	defer s.consumer.Close()

	for {
		select {
		case <-postsUpdated:
			log.Println("Sending new posts to client...")
			posts, err := s.postService.GetPostFeeds(stream.Context())
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

	mediaService := media.NewMediaService(client, cld)

	ka := keepalive.ServerParameters{
		Time:    10 * time.Hour,
		Timeout: 5 * time.Hour,
	}
	s := grpc.NewServer(grpc.KeepaliveParams(ka))

	brokers := []string{"glad-ocelot-13748-eu2-kafka.upstash.io:9092"}
	topic := "post_created"
	username := "Z2xhZC1vY2Vsb3QtMTM3NDgkiJbJsYDFiX7WFPdq0E1rXMVgyy2z-P46ix43a8g"
	password := "MmI0ZmY0MTAtZTU1OS00MjQ0LTkyMmItYjM1MjdhNWY4OThl"

	producer := kafka.NewProducer(brokers, topic, username, password)
	consumer := kafka.NewKafkaConsumer(brokers, topic, "placio", username, password)

	postSvc := posts.NewPostService(client, redisClient, mediaService, producer)
	proto.RegisterPostServiceServer(s, &server{postService: postSvc, producer: producer, consumer: consumer})

	log.Println("gRPC server started on :50051")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

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
