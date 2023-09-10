package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"placio-pkg/kafka"
	socket "placio-realtime/pkg/websocket"
	"placio-realtime/services"

	"github.com/google/uuid"
)

func ServeWs(postService services.PostService, hub *socket.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Failed to establish a WebSocket connection."})
		return
	}
	connection := socket.NewConnection(conn, hub)
	hub.Register <- connection

	response, err := postService.GetPostFeeds(context.Background())
	if err != nil {
		log.Printf("Error getting post feeds: %v", err)
		responseMsg, _ := json.Marshal(ErrorResponse{Message: "Failed to get post feeds.", Status: http.StatusInternalServerError})
		connection.Send <- responseMsg
	} else {
		responseMsg, _ := json.Marshal(response)
		connection.Send <- responseMsg
	}

	brokers := []string{"glad-ocelot-13748-eu2-kafka.upstash.io:9092"}
	topic := "post_created"
	username := "Z2xhZC1vY2Vsb3QtMTM3NDgkiJbJsYDFiX7WFPdq0E1rXMVgyy2z-P46ix43a8g"
	password := "MmI0ZmY0MTAtZTU1OS00MjQ0LTkyMmItYjM1MjdhNWY4OThl"

	consumer := kafka.NewKafkaConsumer(brokers, topic, uuid.New().String(), username, password)

	go func() {
		err := consumer.Start(func(messageValue []byte) error {
			// Process the Kafka message here
			response, err := postService.GetPostFeeds(context.Background())
			if err != nil {
				log.Printf("Error getting post feeds: %v", err)
				return err
			}

			responseMsg, _ := json.Marshal(response)
			connection.Send <- responseMsg
			return nil
		})

		if err != nil {
			log.Printf("Error reading Kafka message: %v", err)
		}
	}()

	go connection.Writer()
	go connection.Reader(postService)
}
