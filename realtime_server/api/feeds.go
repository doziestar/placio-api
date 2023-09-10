package api

import (
	"encoding/json"
	"net/http"
	"placio-realtime/pkg/websocket"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// This can be a mock function for user authentication for now.
func isAuthenticated(r *http.Request) bool {
	// Add your authentication logic here (e.g., JWT token validation).
	// For the sake of this example, we're assuming everyone is authenticated.
	return true
}

// FetchUserFeeds would be the function you use to fetch the user's feeds.
// This is a mock for the sake of this example.
func FetchUserFeeds(userID string) ([]string, error) {
	return []string{"Feed1", "Feed2"}, nil
}

func UserHomeFeedHandler(hub *websocket.Hub, w http.ResponseWriter, r *http.Request) {
	// Check if the user is authenticated.
	if !isAuthenticated(r) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "User is not authenticated."})
		return
	}

	// Here you'd typically fetch the user ID from the authenticated request.
	// For the sake of this example, we're hardcoding a user ID.
	userID := "123"

	feeds, err := FetchUserFeeds(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Failed to fetch feeds."})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}

// func ServeWs(hub *websocket.Hub, w http.ResponseWriter, r *http.Request) {
// 	// Check if the user is authenticated.
// 	if !isAuthenticated(r) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(ErrorResponse{Message: "User is not authenticated."})
// 		return
// 	}

// 	conn, err := websocket.upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(ErrorResponse{Message: "Failed to establish a WebSocket connection."})
// 		return
// 	}

// 	connection := websocket.NewConnection(conn, hub)
// 	hub.Register <- connection

// 	go connection.ReadPump()
// 	go connection.WritePump()
// }

// func watchPostsStream(postService services.PostService, hub *websocket.Hub) {
// 	for {
// 		stream, err := postService.WatchPosts(context.Background())
// 		if err != nil {
// 			if grpcErr, ok := status.FromError(err); ok {
// 				log.Printf("Error watching posts: %s, Code: %s. Retrying in 3 seconds...", grpcErr.Message(), grpcErr.Code())
// 			} else {
// 				log.Printf("Error watching posts: %v. Retrying in 3 seconds...", err)
// 			}

// 			time.Sleep(3 * time.Second)
// 			continue
// 		}

// 		for {
// 			response, err := stream.Recv()
// 			log.Println("Received a post")
// 			if err == io.EOF {
// 				log.Println("Stream ended from server side.")
// 				break
// 			}
// 			if err != nil {
// 				log.Printf("Failed to receive a post: %v", err)
// 				break // break out to outer loop to retry the connection
// 			}
// 			responseMsg, _ := json.Marshal(response)
// 			log.Println("Broadcasting a post")
// 			hub.Broadcast <- responseMsg
// 		}
// 	}
// }