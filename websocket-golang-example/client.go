package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	// Buffer size for message channel
	bufferSize = 100
)

type Client struct {
	ID string

	sender chan Message // Write channel for sending messages to the client

	conn *websocket.Conn
	room *ChatRoom
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		ID:     id,
		sender: make(chan Message, bufferSize),
		conn:   conn,
		room:   nil, // Room will be set when joining
	}
}

func (c *Client) Join(room *ChatRoom) {
	if c.room != nil {
		c.room.unregister <- c
		log.Printf("[DEBUG] Client %s is already in room %s. Leaving and joining room %s\n", c.ID, c.room.RoomId, room.RoomId)
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

// readLoop read messages from the WebSocket connection and broadcasts them to the chat room.
func (c *Client) readLoop() {
	defer func() {
		c.room.unregister <- c
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("[DEBUG] Client disconnected:", c.ID, "Error:", err)
			break
		}

		log.Println("[DEBUG] Received message from", c.ID, ":", string(message))
		c.room.broadcast <- Message{Sender: c.ID, Content: string(message)}
	}
}

func (c *Client) writeLoop() {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case message, ok := <-c.sender:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.CloseMessage()
				return
			}

			if err := c.JsonMessage(message); err != nil {
				log.Printf("[DEBUG] Error writing message to client: %s, error: %v\n", c.ID, err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.PingMessage(); err != nil {
				log.Printf("[DEBUG] Ping error, closing client: %s, error: %v\n", c.ID, err)
				return
			}
		}
	}
}

func (c *Client) JsonMessage(message Message) error {
	return c.conn.WriteJSON(message)
}

func (c *Client) PingMessage() error {
	return c.conn.WriteMessage(websocket.PingMessage, nil)
}

func (c *Client) CloseMessage() error {
	return c.conn.WriteMessage(websocket.CloseMessage, []byte{})
}
