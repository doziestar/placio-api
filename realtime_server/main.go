package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"placio-pkg/grpc/proto"
	"placio-pkg/middleware"
	"placio-realtime/api"
	"placio-realtime/pkg/websocket"
	"placio-realtime/services"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	serverAddress = "placio-backend:50051"
	serverPort    = ":7080"
	shutdownDelay = 5 * time.Second
)

func main() {
	r := setupRouter()
	startHub(r)
	startServer(r)
}

func setupRouter() *mux.Router {
	log.Println("Setting up router...")
	r := mux.NewRouter()
	return r
}

func createPostServiceClient(conn *grpc.ClientConn) services.PostService {
	log.Println("Creating post service client...")

	c := proto.NewPostServiceClient(conn)

	return services.NewPostService(c)
}

func startHub(r *mux.Router) {
	log.Println("Starting websocket hub...")
	hub := websocket.NewHub()
	go hub.Run()
	//go watchPostsStream(postService, hub)
}

func startServer(r *mux.Router) {
	log.Println("Starting server...")
	http.Handle("/", r)

	srv := &http.Server{
		Addr:    serverPort,
		Handler: r,
	}

	// Create connection to gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close() // Close connection when main() exits
	postServiceClient := createPostServiceClient(conn)
	r.HandleFunc("/home-feeds", middleware.EnsureValidWebSocketToken(func(w http.ResponseWriter, r *http.Request) {
		hub := websocket.NewHub()
		api.ServeWs(postServiceClient, hub, w, r)
	}))

	// Create channel to listen for interrupt or terminate signal from OS
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("-------------------------------------------------")
		log.Printf("Websocket Hub started on port %s\n", serverPort)
		log.Println("-------------------------------------------------")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v\n", err)
		}
	}()

	<-stop

	// Graceful Shutdown
	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownDelay)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server gracefully stopped.")
}
