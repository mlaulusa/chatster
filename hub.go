package main

type Hub struct {
	clients map[*Client] bool
	register chan *Client
	broadcast chan []byte
}

func CreateHub () *Hub {
	return &Hub{
		clients: make(map[*Client] bool),
		register: make(chan *Client),
		broadcast: make(chan []byte),
	}

}

func (hub *Hub) Run () {
	for {
		select {
			case client := <- hub.register:
				hub.clients[client] = true

			case message := <- hub.broadcast:
					for client := range hub.clients {
						client.send <- message

					}
		}
	}
}
