package homefeeds

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"placio-api/events/kafka"
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
	producer     *kafka.Producer
	consumer     *kafka.KafkaConsumer
}

func NewHomeFeedsHandler(postService posts.PostService, producer *kafka.Producer, consumer *kafka.KafkaConsumer, placeService places.PlaceService) *HomeFeedsHandler {
	return &HomeFeedsHandler{postService: postService, placeService: placeService, producer: producer, consumer: consumer}
}

func (s *HomeFeedsHandler) HandleHomeFeeds(ctx context.Context, ws *websocket.Conn) {
	log.Println("Handling Home Feeds", ctx.Value("user"))
	clients[ws] = true

	type Message struct {
		Post           []*ent.Post           `json:"post"`
		Place          []*ent.Place          `json:"place"`
		PlaceInventory []*ent.PlaceInventory `json:"place_inventory"`
	}

	//// Subscribe to post:created event
	//log.Println("Subscribing to post:created event")
	//subscription := s.eventBus.Subscribe("post:created", func(post *ent.Post) {
	//	err := ws.WriteJSON(post)
	//	if err != nil {
	//		log.Printf("error writing to websocket: %v", err)
	//	}
	//})
	//
	//s.consumer.Start()

	defer func() {
		ws.Close()
		delete(clients, ws)
		//s.eventBus.Unsubscribe("post:created", subscription)
	}()

	log.Println("Subscribed to post:created event")

	// Fetch all posts
	log.Println("Fetching all posts")
	posts, err := s.postService.GetPostFeeds(ctx)
	if err != nil {
		log.Printf("error fetching posts: %v", err)
		return
	}

	type Error struct {
		Message string `json:"message"`
	}

	log.Println("Fetching all places", ctx.Value("user"))
	//placeData, _, err := s.placeService.GetPlaces(ctx, nil, "", 10)
	//if err != nil {
	//	log.Printf("error fetching places: %v", err)
	//	errMessage := Error{Message: "error writing to websocket: " + err.Error()}
	//
	//	// Convert the Error instance to JSON
	//	jsonError, jsonErr := json.Marshal(errMessage)
	//	if jsonErr != nil {
	//		log.Printf("error converting error message to JSON: %v", jsonErr)
	//		return
	//	}
	//
	//	// Send the JSON error message to the client
	//	ws.WriteMessage(websocket.TextMessage, jsonError)
	//	return
	//}

	message := Message{
		Post: posts,
		//Place: placeData,
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
	log.Println("Sending posts to the connected client")
	err = ws.WriteMessage(websocket.TextMessage, jsonPosts)
	if err != nil {
		// Create an Error instance
		errMessage := Error{Message: "error writing to websocket: " + err.Error()}

		// Convert the Error instance to JSON
		jsonError, jsonErr := json.Marshal(errMessage)
		if jsonErr != nil {
			log.Printf("error converting error message to JSON: %v", jsonErr)
			return
		}

		// Send the JSON error message to the client
		ws.WriteMessage(websocket.TextMessage, jsonError)

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
		type messageReceived struct {
			Refresh bool `json:"refresh"`
		}

		var messageReceivedData messageReceived
		err = json.Unmarshal(msg, &messageReceivedData)
		if err != nil {
			log.Printf("error unmarshalling message: %v", err)
			break
		}

		if messageReceivedData.Refresh {
			// Fetch all posts
			log.Println("Fetching all posts")
			posts, err := s.postService.GetPostFeeds(ctx)
			if err != nil {
				log.Printf("error fetching posts: %v", err)
				return
			}

			type Error struct {
				Message string `json:"message"`
			}

			log.Println("Fetching all places", ctx.Value("user"))
			//placeData, _, err := s.placeService.GetPlaces(ctx, nil, "", 10)
			//if err != nil {
			//	log.Printf("error fetching places: %v", err)
			//	errMessage := Error{Message: "error writing to websocket: " + err.Error()}
			//
			//	// Convert the Error instance to JSON
			//	jsonError, jsonErr := json.Marshal(errMessage)
			//	if jsonErr != nil {
			//		log.Printf("error converting error message to JSON: %v", jsonErr)
			//		return
			//	}
			//
			//	// Send the JSON error message to the client
			//	ws.WriteMessage(websocket.TextMessage, jsonError)
			//	return
			//}

			message := Message{
				Post: posts,
				//Place: placeData,
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
			log.Println("Sending posts to the connected client")
			err = ws.WriteMessage(websocket.TextMessage, jsonPosts)
			if err != nil {
				// Create an Error instance
				errMessage := Error{Message: "error writing to websocket: " + err.Error()}

				// Convert the Error instance to JSON
				jsonError, jsonErr := json.Marshal(errMessage)
				if jsonErr != nil {
					log.Printf("error converting error message to JSON: %v", jsonErr)
					return
				}

				// Send the JSON error message to the client
				ws.WriteMessage(websocket.TextMessage, jsonError)

				return
			}
		}

		// Broadcast message to all connected clients
		log.Println("Broadcasting message to all connected clients")
		//s.BroadcastMessage(msg)
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
