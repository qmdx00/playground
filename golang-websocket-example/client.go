package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID string

	conn *websocket.Conn
	room *ChatRoom
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		ID:   id,
		conn: conn,
		room: nil, // Room will be set when joining
	}
}

func (c *Client) Join(room *ChatRoom) {
	if c.room != nil {
		c.room.unregister <- c
		log.Printf("[DEBUG] Client %s is already in room %s. Leaving and joining room %s\n", c.ID, c.room.roomId, room.roomId)
	}

	c.room = room
	c.room.register <- c
}

func (c *Client) Leave() {
	if c.room != nil {
		c.room.unregister <- c
		c.room = nil
	}
}

func (c *Client) Close() error {
	log.Println("[DEBUG] Closing connection for client:", c.ID)
	if err := c.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (c *Client) WriteMessage(message string) error {
	return c.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

// readLoop read messages from the WebSocket connection and broadcasts them to the chat room.
func (c *Client) readLoop() {
	defer func() {
		c.room.unregister <- c
		_ = c.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("[DEBUG] Client disconnected:", c.ID)
			break
		}

		log.Println("[DEBUG] Received message from", c.ID, ":", string(message))
		c.room.broadcast <- Message{Sender: c.ID, Content: string(message)}
	}
}
