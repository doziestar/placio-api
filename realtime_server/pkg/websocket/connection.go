package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Modify this to set up your CORS restrictions
	},
}

type Connection struct {
	ws   *websocket.Conn
	send chan []byte
	hub  *Hub
}

func NewConnection(ws *websocket.Conn, hub *Hub) *Connection {
	return &Connection{ws: ws, send: make(chan []byte), hub: hub}
}

func (c *Connection) Reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.hub.Broadcast <- message
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
