package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients []*Client
	register chan *Client
	broadcast chan *Broadcast
}

func CreateHub () *Hub {
	return &Hub{
		clients: make([]*Client, 0),
		register: make(chan *Client),
		broadcast: make(chan *Broadcast),
	}

}

func (hub *Hub) Run () {
	for {
		select {
			case client := <- hub.register:
				hub.clients = append(hub.clients, client)
				writer, err := client.connection.NextWriter(websocket.TextMessage)

				if err != nil {
					log.Print(err.Error())
				}

				i, err := writer.Write([]byte("Welcome!"))

				if err != nil {
					log.Print(i)
					log.Print(err.Error())
				}

			case broadcast := <- hub.broadcast:
				for _, client := range hub.clients {
					if client != broadcast.from {
						client.send <- broadcast.message
					}
				}

		}
	}
}
