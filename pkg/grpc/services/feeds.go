package services

import (
	"context"
	"log"
	"placio-app/domains/posts"
	"placio-app/ent"
	"placio-pkg/grpc/proto"
	"placio-pkg/kafka"
	"sync"
)

type FeedService interface {
	RefreshPost(ctx context.Context, req *proto.RefreshPostRequest) (*proto.GetPostFeedsResponse, error)
	GetPostFeeds(ctx context.Context, req *proto.GetPostFeedsRequest) (*proto.GetPostFeedsResponse, error)
	WatchPosts(stream proto.PostService_WatchPostsServer) error
}

type Server struct {
	proto.UnimplementedPostServiceServer
	Producer    *kafka.Producer
	Consumer    *kafka.KafkaConsumer
	PostService posts.PostService
}

func NewServer(postService posts.PostService, producer *kafka.Producer, consumer *kafka.KafkaConsumer) proto.PostServiceServer {
	return &Server{PostService: postService, Producer: producer, Consumer: consumer}
}

func convertToPbPost(p *ent.Post) *proto.Post {

	if p == nil {
		return nil
	}
	var user *proto.Post_User
	if p.Edges.User != nil {
		user = &proto.Post_User{
			Id:       p.Edges.User.ID,
			Username: p.Edges.User.Username,
			Name:     p.Edges.User.Name,
			Picture:  p.Edges.User.Picture,
		}
	}

	var media []*proto.Post_Media
	for _, m := range p.Edges.Medias {
		if m != nil {
			media = append(media, &proto.Post_Media{
				Id:        m.ID,
				URL:       m.URL,
				MediaType: m.MediaType,
			})
		}
	}

	var pbComments []*proto.Post_Comment
	for _, c := range p.Edges.Comments {
		if c != nil {
			pbComments = append(pbComments, &proto.Post_Comment{
				Id:        c.ID,
				Content:   c.Content,
				CreatedAt: c.CreatedAt.String(),
				UpdatedAt: c.UpdatedAt.String(),
				Edges: &proto.Post_CommentEdge{
					User: &proto.Post_User{
						Id:       c.Edges.User.ID,
						Username: c.Edges.User.Username,
						Name:     c.Edges.User.Name,
						Picture:  c.Edges.User.Picture,
					},
				},
			})
		}
	}

	return &proto.Post{
		Id:           p.ID,
		Content:      p.Content,
		CreatedAt:    p.CreatedAt.String(),
		UpdatedAt:    p.UpdatedAt.String(),
		LikeCount:    int64(p.LikeCount),
		CommentCount: int64(p.CommentCount),
		LikedByMe:    p.LikedByMe,
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
			User:     user,
			Comments: pbComments,
			Media:    media,
		},
	}
}

func (s *Server) RefreshPost(ctx context.Context, req *proto.RefreshPostRequest) (*proto.GetPostFeedsResponse, error) {
	posts, err := s.PostService.GetPostFeeds(ctx)
	if err != nil {
		log.Printf("Error getting post feeds: %v", err)
		return nil, err
	}

	var pbPosts []*proto.Post

	var wg sync.WaitGroup

	wg.Add(len(posts))
	for _, p := range posts {
		go func(p *ent.Post) {
			pbPosts = append(pbPosts, convertToPbPost(p))
			wg.Done()
		}(p)
	}
	wg.Wait()

	return &proto.GetPostFeedsResponse{Posts: pbPosts}, nil
}

func (s *Server) GetPostFeeds(ctx context.Context, req *proto.GetPostFeedsRequest) (*proto.GetPostFeedsResponse, error) {
	feedPosts, err := s.PostService.GetPostFeeds(ctx)
	if err != nil {
		log.Printf("Error getting post feeds: %v", err)
		return nil, err
	}

	var pbPosts []*proto.Post
	var wg sync.WaitGroup

	wg.Add(len(feedPosts))
	for _, p := range feedPosts {
		go func(p *ent.Post) {
			pbPosts = append(pbPosts, convertToPbPost(p))
			wg.Done()
		}(p)
	}
	wg.Wait()

	return &proto.GetPostFeedsResponse{Posts: pbPosts}, nil
}

func (s *Server) WatchPosts(stream proto.PostService_WatchPostsServer) error {
	postsUpdated := make(chan bool, 100)

	log.Println("Client connected to WatchPosts stream.")
	defer s.Consumer.Close()

	for {
		select {
		case <-postsUpdated:
			log.Println("Sending new posts to client...")
			posts, err := s.PostService.GetPostFeeds(stream.Context())
			if err != nil {
				return err
			}

			var pbPosts []*proto.Post
			var wg sync.WaitGroup

			wg.Add(len(posts))
			for _, p := range posts {
				go func(p *ent.Post) {
					pbPosts = append(pbPosts, convertToPbPost(p))
					wg.Done()
				}(p)
			}
			wg.Wait()

			if err := stream.Send(&proto.GetPostFeedsResponse{Posts: pbPosts}); err != nil {
				return err
			}

		case <-stream.Context().Done():
			log.Println("Client disconnected from WatchPosts stream.")
			return stream.Context().Err()
		}
	}
}
