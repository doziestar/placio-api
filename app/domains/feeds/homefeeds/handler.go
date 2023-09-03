package homefeeds

import (
	"context"
	"encoding/json"
	"github.com/asaskevich/EventBus"
	"log"
	"placio-app/domains/posts"
	"placio-app/ent"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)

type IHomeFeedsHandler interface {
	HandleHomeFeeds()
	BroadcastMessage(msg []byte)
}

type HomeFeedsHandler struct {
	ws          *websocket.Conn
	postService posts.PostService
	eventBus    EventBus.Bus
}

func NewHomeFeedsHandler(postService posts.PostService, eventBus EventBus.Bus) *HomeFeedsHandler {
	return &HomeFeedsHandler{postService: postService, eventBus: eventBus}
}

func (s *HomeFeedsHandler) HandleHomeFeeds() {
	log.Println("Handling Home Feeds")
	clients[s.ws] = true

	// Subscribe to post:created event
	log.Println("Subscribing to post:created event")
	subscription := s.eventBus.Subscribe("post:created", func(post *ent.Post) {
		err := s.ws.WriteJSON(post)
		if err != nil {
			log.Printf("error writing to websocket: %v", err)
		}
	})

	defer func() {
		s.ws.Close()
		delete(clients, s.ws)
		s.eventBus.Unsubscribe("post:created", subscription)
	}()

	log.Println("Subscribed to post:created event")

	// Fetch all posts
	log.Println("Fetching all posts")
	ctx := context.Background()
	posts, err := s.postService.GetPostFeeds(ctx)
	if err != nil {
		log.Printf("error fetching posts: %v", err)
		return
	}
	log.Println("Fetched all posts")

	// Convert posts to JSON
	jsonPosts, err := json.Marshal(posts)
	if err != nil {
		log.Printf("error converting posts to JSON: %v", err)
		return
	}
	log.Println("Converted posts to JSON")

	// Send posts to the connected client
	log.Println("Sending posts to the connected client")
	err = s.ws.WriteJSON(jsonPosts)
	if err != nil {
		log.Printf("error writing to websocket: %v", err)
		return
	}
	log.Println("Sent posts to the connected client")

	for {
		// Read message from browser
		mt, msg, err := s.ws.ReadMessage()
		if err != nil {
			log.Printf("error reading websocket message: %v", err)
			break
		}

		// Check if the client sent a close message
		if mt == websocket.CloseMessage {
			break
		}

		log.Printf("Received: %s", msg)

		// Broadcast message to all connected clients
		log.Println("Broadcasting message to all connected clients")
		s.BroadcastMessage(msg)
	}
}

func (s *HomeFeedsHandler) BroadcastMessage(msg []byte) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("error writing to websocket: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
