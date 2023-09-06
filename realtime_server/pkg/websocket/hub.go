package websocket

import "log"

type Hub struct {
	Connections map[*Connection]bool
	Broadcast   chan []byte
	Register    chan *Connection
	Unregister  chan *Connection
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:   make(chan []byte),
		Register:    make(chan *Connection),
		Unregister:  make(chan *Connection),
		Connections: make(map[*Connection]bool),
	}
}

func (h *Hub) Run() {
	log.Println("Hub Run")
	for {
		select {
		case conn := <-h.Register:
			h.Connections[conn] = true
		case conn := <-h.Unregister:
			if _, ok := h.Connections[conn]; ok {
				delete(h.Connections, conn)
				close(conn.Send)
			}
		case message := <-h.Broadcast:
			for conn := range h.Connections {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(h.Connections, conn)
				}
			}
		}
	}
}
