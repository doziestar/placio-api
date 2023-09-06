package main

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"placio-api/grpc/proto"
	"placio-realtime/api"
	"placio-realtime/pkg/websocket"
	"placio-realtime/services"
)

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

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":7080", nil))
}
