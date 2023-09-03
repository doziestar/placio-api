package homefeeds

import (
	"context"
	"encoding/json"
	"log"
	"placio-app/domains/posts"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)

func HandleHomeFeeds(ws *websocket.Conn, postService posts.PostService) {
	defer ws.Close()
	clients[ws] = true

	// Fetch all posts
	ctx := context.Background()
	posts, err := postService.GetPostFeeds(ctx)
	if err != nil {
		log.Printf("error fetching posts: %v", err)
		return
	}

	// Convert posts to JSON
	jsonPosts, err := json.Marshal(posts)
	if err != nil {
		log.Printf("error converting posts to JSON: %v", err)
		return
	}

	// Send posts to the connected client
	err = ws.WriteMessage(websocket.TextMessage, jsonPosts)
	if err != nil {
		log.Printf("error writing to websocket: %v", err)
		delete(clients, ws)
		return
	}

	for {
		// Read message from browser
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error reading websocket message: %v", err)
			delete(clients, ws)
			break
		}

		log.Printf("Received: %s", msg)

		// Broadcast message to all connected clients
		BroadcastMessage(msg)
	}
}

func BroadcastMessage(msg []byte) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("error writing to websocket: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
