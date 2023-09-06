package feeds

import (
	"github.com/getsentry/sentry-go"
	"log"
	"net/http"
	"placio-api/events/kafka"
	"placio-app/domains/feeds/chats"
	"placio-app/domains/feeds/homefeeds"
	"placio-app/domains/places"
	"placio-app/domains/posts"
	"placio-pkg/middleware"
	"sync"

	"github.com/gorilla/websocket"
)

type IWebSocketServer interface {
	AddClient(client *websocket.Conn)
	RemoveClient(client *websocket.Conn)
	BroadcastMessage(message []byte)
	RoomBroadcastMessage(room string, message []byte)
	RoomAddClient(room string, client *websocket.Conn)
	RoomRemoveClient(room string, client *websocket.Conn)
	RUN(port string) error
	HandleConnections(w http.ResponseWriter, r *http.Request)
}

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	producer  *kafka.Producer
	consumer  *kafka.KafkaConsumer
	clientsMu sync.Mutex
	upgrader  websocket.Upgrader
	homefeeds homefeeds.IHomeFeedsHandler
}

func NewWebSocketServer(producer *kafka.Producer, consumer *kafka.KafkaConsumer, postService posts.PostService, placeService places.PlaceService) *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		homefeeds: homefeeds.NewHomeFeedsHandler(postService, producer, consumer, placeService),
	}
}

func (s *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	err := middleware.EnsureValidWebSocketToken(w, r)
	if err != nil {
		log.Println(err)
		sentry.CaptureException(err)
		return
	}

	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		sentry.CaptureException(err)
		return
	}
	defer ws.Close()

	// Handle different functionalities based on the path
	switch r.URL.Path {
	case "/chat":
		chats.HandleChat(ws)
	case "/home-feeds":
		s.homefeeds.HandleHomeFeeds(r.Context(), ws)
	default:
		log.Printf("Unknown path: %s", r.URL.Path)
	}
}

func (s *WebSocketServer) AddClient(client *websocket.Conn) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	s.clients[client] = true
}

func (s *WebSocketServer) RemoveClient(client *websocket.Conn) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	delete(s.clients, client)
}

func (s *WebSocketServer) BroadcastMessage(message []byte) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	for client := range s.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("error sending message to client %v: %v", client.RemoteAddr(), err)
			delete(s.clients, client)
			if err := client.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				log.Printf("error closing client %v: %v", client.RemoteAddr(), err)
			}
			if err := client.Close(); err != nil {
				log.Printf("error closing client %v: %v", client.RemoteAddr(), err)
			}
		}
	}
}

func (s *WebSocketServer) RoomBroadcastMessage(room string, message []byte) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	for client := range s.clients {
		if client.Subprotocol() == room {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("error sending message to client %v: %v", client.RemoteAddr(), err)
				delete(s.clients, client)
				if err := client.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
					log.Printf("error closing client %v: %v", client.RemoteAddr(), err)
				}
				if err := client.Close(); err != nil {
					log.Printf("error closing client %v: %v", client.RemoteAddr(), err)
				}
			}
		}
	}
}

func (s *WebSocketServer) RoomAddClient(room string, client *websocket.Conn) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	s.clients[client] = true
}

func (s *WebSocketServer) RoomRemoveClient(room string, client *websocket.Conn) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	delete(s.clients, client)
}

func (s *WebSocketServer) RUN(addr string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalf("error upgrading request to a websocket:: %v", err)
			return
		}
		defer ws.Close()

		s.AddClient(ws)

		for {
			// Read message from browser
			_, msg, err := ws.ReadMessage()
			if err != nil {
				s.RemoveClient(ws)
				log.Printf("error reading websocket message: %v", err)
				break
			}

			log.Printf("Received: %s", msg)

			// Broadcast message to all connected clients
			s.BroadcastMessage(msg)
		}
	})

	// Start the server
	log.Printf("WebSocket server started on: %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
		return err
	}

	return nil
}
