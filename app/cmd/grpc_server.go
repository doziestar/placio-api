package cmd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"placio-pkg/grpc/proto"
	"placio-pkg/grpc/services"
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
	server := services.NewServer(postSvc, producer, consumer)
	proto.RegisterPostServiceServer(s, server)

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
