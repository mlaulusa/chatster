package socket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Room struct {
	clients []*Client
	Register chan *Client
	broadcast chan *Broadcast
}

func NewRoom () *Room {
	return &Room{
		clients: make([]*Client, 0),
		Register: make(chan *Client),
		broadcast: make(chan *Broadcast),
	}
}

func (room *Room) Broadcast (broadcast Broadcast) {
	for _, client := range room.clients {
		if client != broadcast.from {
			client.send <- broadcast.message
		}
	}
}

func (room *Room) Run () {
	for {
		select {
		case client := <- room.Register:
			room.clients = append(room.clients, client)
			writer, err := client.connection.NextWriter(websocket.TextMessage)

			if err != nil {
				log.Print(err.Error())
			}

			i, err := writer.Write([]byte("Welcome!"))

			if err != nil {
				log.Print(i)
				log.Print(err.Error())
			}

			if err := writer.Close(); err != nil {
				return
			}

		case broadcast := <- room.broadcast:
			for _, client := range room.clients {
				if client != broadcast.from {
					client.send <- broadcast.message
				}
			}

		}
	}
}
