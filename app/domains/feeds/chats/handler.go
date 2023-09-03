package chats

import (
	"github.com/gorilla/websocket"
	"log"
)

var clients = make(map[*websocket.Conn]bool)

func HandleChat(ws *websocket.Conn) {
	log.Println("Handling chat")
	defer ws.Close()

	for {
		// Read message from browser
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error reading websocket message: %v", err)
			break
		}

		log.Printf("Received: %s", msg)

		// Broadcast message to all connected clients
		// Note: You need to implement the BroadcastMessage method
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
