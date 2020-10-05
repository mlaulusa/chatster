package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	clients = make(map[*websocket.Conn] bool)

	upgrader = websocket.Upgrader{
		CheckOrigin: func (r *http.Request) bool {
			return true
		},
	}
)

func upgrade (w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, nil)
}

func handleWebSocket (hub *Hub) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {

		ws, err := upgrade(w, r)

		if err != nil {
			log.Fatal(err)
			return
		}

		client := NewClient(hub, ws)

		hub.register <- client

		go client.read()
		go client.write()

	}
}

func main () {
	fs := http.FileServer(http.Dir("./public/build"))

	hub := CreateHub()

	go hub.Run()

	http.Handle("/", fs)

	http.HandleFunc("/ws", handleWebSocket(hub))

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
