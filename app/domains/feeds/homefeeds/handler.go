package homefeeds

import (
	"context"
	"encoding/json"
	"github.com/asaskevich/EventBus"
	"github.com/gorilla/websocket"
	"log"
	"placio-app/domains/places"
	"placio-app/domains/posts"
	"placio-app/ent"
)

var clients = make(map[*websocket.Conn]bool)

type IHomeFeedsHandler interface {
	HandleHomeFeeds(ctx context.Context, ws *websocket.Conn)
	BroadcastMessage(msg []byte)
}

type HomeFeedsHandler struct {
	postService  posts.PostService
	placeService places.PlaceService
	eventBus     EventBus.Bus
}

func NewHomeFeedsHandler(postService posts.PostService, eventBus EventBus.Bus, placeService places.PlaceService) *HomeFeedsHandler {
	return &HomeFeedsHandler{postService: postService, placeService: placeService, eventBus: eventBus}
}

func (s *HomeFeedsHandler) HandleHomeFeeds(ctx context.Context, ws *websocket.Conn) {
	log.Println("Handling Home Feeds", ctx.Value("user"))
	clients[ws] = true

	type Message struct {
		Post           []*ent.Post           `json:"post"`
		Place          []*ent.Place          `json:"place"`
		PlaceInventory []*ent.PlaceInventory `json:"place_inventory"`
	}

	// Subscribe to post:created event
	log.Println("Subscribing to post:created event")
	subscription := s.eventBus.Subscribe("post:created", func(post *ent.Post) {
		err := ws.WriteJSON(post)
		if err != nil {
			log.Printf("error writing to websocket: %v", err)
		}
	})

	defer func() {
		ws.Close()
		delete(clients, ws)
		s.eventBus.Unsubscribe("post:created", subscription)
	}()

	log.Println("Subscribed to post:created event")

	// Fetch all posts
	log.Println("Fetching all posts")
	posts, err := s.postService.GetPostFeeds(ctx)
	if err != nil {
		log.Printf("error fetching posts: %v", err)
		return
	}

	//log.Println("Fetching all places", ctx.Value("user"))
	//places, _, err := s.placeService.GetPlaces(ctx, nil, "", 10)
	//if err != nil {
	//	log.Printf("error fetching places: %v", err)
	//	return
	//}

	message := Message{
		Post: posts,
		//Place: places,
	}
	log.Println("Fetched all posts")

	// Convert posts to JSON
	jsonPosts, err := json.Marshal(message)
	if err != nil {
		log.Printf("error converting posts to JSON: %v", err)
		return
	}
	log.Println("Converted posts to JSON")

	// Send posts to the connected client
	// Send posts to the connected client
	log.Println("Sending posts to the connected client")
	err = ws.WriteMessage(websocket.TextMessage, jsonPosts)
	if err != nil {
		log.Printf("error writing to websocket: %v", err)
		return
	}
	log.Println("Sent posts to the connected client")

	for {
		// Read message from browser
		mt, msg, err := ws.ReadMessage()
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
