package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/mlaulusa/chatster/model"
	"github.com/mlaulusa/chatster/route"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, nil)
}

func handleMessages(hub *Hub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

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

func main() {

	defer model.Close()

	react := http.FileServer(http.Dir("./static/react"))

	hub := CreateHub()

	go hub.Run()

	// router.HandleFunc("/room/{room}", handleWebSocket(hub))
	// router.HandleFunc("/ws", handleWebSocket(hub))

	http.Handle("/room", route.GetRoomRouter())
	// http.HandleFunc("/ws", handleWebSocket(hub))

	http.Handle("/react", react)
	http.Handle("/", react)

	err := http.ListenAndServe(":3000", nil)

	// server := &http.Server{
	// 	Addr:              "localhost:3000",
	// 	Handler:           router,
	// 	ReadTimeout:       15 * time.Second,
	// 	WriteTimeout:      15 * time.Second,
	// }

	// err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
