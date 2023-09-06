package websocket

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"placio-realtime/services"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Modify this to set up your CORS restrictions
	},
}

type WsMessage struct {
	Type string `json:"type"`
}

type Connection struct {
	ws   *websocket.Conn
	send chan []byte
	hub  *Hub
}

func NewConnection(ws *websocket.Conn, hub *Hub) *Connection {
	return &Connection{ws: ws, send: make(chan []byte), hub: hub}
}

func (c *Connection) Reader(postService services.PostService) {
	for {
		_, p, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		var message WsMessage
		if err := json.Unmarshal(p, &message); err != nil {
			log.Printf("Error decoding message: %v", err)
			continue
		}

		var responseMsg []byte

		switch message.Type {
		case "getFeeds":
			response, err := postService.GetPostFeeds(context.Background())
			if err != nil {
				log.Printf("Error getting post feeds: %v", err)
			} else {
				responseMsg, _ = json.Marshal(response) // serialize response to JSON
			}
		case "refresh":
			response, err := postService.RefreshPostFeeds(context.Background())
			if err != nil {
				log.Printf("Error refreshing post feeds: %v", err)
			} else {
				responseMsg, _ = json.Marshal(response) // serialize response to JSON
			}
		default:
			log.Printf("Unknown message type: %s", message.Type)
		}

		if responseMsg != nil {
			c.hub.Broadcast <- responseMsg
		}
	}
	c.ws.Close()
}

func (c *Connection) Writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
