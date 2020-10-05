package main

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	heartbeatTick = 10 * time.Second

	writeDeadline = 10 * time.Second

	writeChannelSize = 256
)

var (
	newline = []byte{'\n'}

	space = []byte{' '}

)

type Client struct {
	hub *Hub
	connection *websocket.Conn
	send chan []byte
}

func NewClient (hub *Hub, connection *websocket.Conn) *Client {
	connection.SetPongHandler(func (appData string) error {
		log.Print("Received pong message")
		return nil
	})

	return &Client{
		hub: hub,
		connection: connection,
		send: make(chan []byte, writeChannelSize),
	}
}

func (c *Client) read () {
	for {
		_, message, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Print(string(message))

		c.hub.broadcast <- bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

	}
}

func (c *Client) write () {

	heartbeat := time.NewTicker(heartbeatTick)

	defer func () {
		heartbeat.Stop()
		c.connection.Close()
	}()

	for {
		select {

		case message, ok := <- c.send:
			c.connection.SetWriteDeadline(time.Now().Add(writeDeadline))

			if !ok {

				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := c.connection.NextWriter(websocket.TextMessage)

			if err != nil {
				return
			}

			writer.Write(message)

			for bufferedMessage := range c.send {
				writer.Write(newline)
				writer.Write(bufferedMessage)
			}

			if err := writer.Close(); err != nil {
				return
			}

			case <- heartbeat.C:
				log.Print("Sending ping message")
				c.connection.SetWriteDeadline(time.Now().Add(writeDeadline))
				if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
		}
	}

}
