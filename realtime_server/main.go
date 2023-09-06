package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"placio-realtime/api"
	"placio-realtime/pkg/websocket"
)

func main() {
	r := mux.NewRouter()
	hub := websocket.NewHub()

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		api.ServeWs(hub, w, r)
	})

	go hub.Run()

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":7071", nil))
}
