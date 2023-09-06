package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net/http"
	"placio-api/grpc/proto"
	"placio-realtime/api"
	"placio-realtime/pkg/websocket"
	"placio-realtime/services"
	"time"
)

func watchPostsStream(postService services.PostService, hub *websocket.Hub) {
	for {
		stream, err := postService.WatchPosts(context.Background())
		if err != nil {
			if grpcErr, ok := status.FromError(err); ok {
				log.Printf("Error watching posts: %s, Code: %s. Retrying in 3 seconds...", grpcErr.Message(), grpcErr.Code())
			} else {
				log.Printf("Error watching posts: %v. Retrying in 3 seconds...", err)
			}

			time.Sleep(3 * time.Second)
			continue
		}

		for {
			response, err := stream.Recv()
			log.Println("Received a post")
			if err == io.EOF {
				log.Println("Stream ended from server side.")
				break
			}
			if err != nil {
				log.Printf("Failed to receive a post: %v", err)
				break // break out to outer loop to retry the connection
			}
			responseMsg, _ := json.Marshal(response)
			log.Println("Broadcasting a post")
			hub.Broadcast <- responseMsg
		}
	}
}

func main() {
	r := mux.NewRouter()
	hub := websocket.NewHub()

	conn, err := grpc.Dial("placio-backend:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	c := proto.NewPostServiceClient(conn)

	postService := services.NewPostService(c)

	r.HandleFunc("/home-feeds", func(w http.ResponseWriter, r *http.Request) {
		api.ServeWs(postService, hub, w, r)
	})

	go hub.Run()

	go watchPostsStream(postService, hub)

	http.Handle("/", r)

	log.Println("---------------------------")
	log.Println("Hub Run: Port 7080")
	log.Println("---------------------------")
	log.Fatal(http.ListenAndServe(":7080", nil))
}
