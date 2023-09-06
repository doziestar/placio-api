package api

import (
	"encoding/json"
	"log"
	"net/http"
	socket "placio-realtime/pkg/websocket"
	"placio-realtime/services"
)

func ServeWs(postService services.PostService, hub *socket.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Failed to establish a WebSocket connection."})
		return
	}
	connection := socket.NewConnection(conn, hub)
	hub.Register <- connection

	response, err := postService.GetPostFeeds(r.Context())
	if err != nil {
		log.Printf("Error getting post feeds: %v", err)
	} else {
		responseMsg, _ := json.Marshal(response)
		connection.Send <- responseMsg
	}

	go connection.Writer()
	go connection.Reader(postService)
}
