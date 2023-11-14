package hub

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte

	Name string
}

func (c *Client) WaitingForMessage() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(1 * time.Minute)); return nil })
	for {
		messageType, message, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		if messageType == websocket.PingMessage {
			continue
		}
		fmt.Printf("%s has send message: %s", c.Name, message)
		c.Hub.Broadcast <- message
	}
}
