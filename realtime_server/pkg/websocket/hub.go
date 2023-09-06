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
	for {
		log.Println("---------------------------")
		log.Println("Hub Run: Port 7080")
		log.Println("---------------------------")
		select {
		case conn := <-h.Register:
			h.Connections[conn] = true
		case conn := <-h.Unregister:
			if _, ok := h.Connections[conn]; ok {
				delete(h.Connections, conn)
				close(conn.send)
			}
		case message := <-h.Broadcast:
			for conn := range h.Connections {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(h.Connections, conn)
				}
			}
		}
	}
}
