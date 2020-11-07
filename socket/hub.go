package socket

type Hub struct {
	//clients []*Client
	//Register chan *Client
	//broadcast chan *Broadcast
	rooms map[string]*Room
}

func CreateHub () *Hub {
	return &Hub{
		//clients: make([]*Client, 0),
		//Register: make(chan *Client),
		//broadcast: make(chan *Broadcast),
		rooms: make(map[string]*Room),
	}
}

func (hub *Hub) GetRoom (name string) *Room {
	if room, ok := hub.rooms[name]; ok {
		return room
	}

	room := NewRoom()

	go room.Run()

	hub.rooms[name] = room

	return room
}

func (hub *Hub) RemoveRoom (name string) {
	delete(hub.rooms, name)
}

//func (hub *Hub) Run () {
//	for {
//		select {
//			case client := <- hub.Register:
//				hub.clients = append(hub.clients, client)
//				writer, err := client.connection.NextWriter(websocket.TextMessage)
//
//				if err != nil {
//					log.Print(err.Error())
//				}
//
//				i, err := writer.Write([]byte("Welcome!"))
//
//				if err != nil {
//					log.Print(i)
//					log.Print(err.Error())
//				}
//
//				if err := writer.Close(); err != nil {
//					return
//				}
//
//			case broadcast := <- hub.broadcast:
//				for _, client := range hub.clients {
//					if client != broadcast.from {
//						client.send <- broadcast.message
//					}
//				}
//
//		}
//	}
//}
