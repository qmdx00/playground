package main

import (
	"fmt"
	"log"
	"sync"
)

const (
	SystemSender = "SYSTEM"
)

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

type ChatRoom struct {
	RoomId string

	register   chan *Client
	unregister chan *Client
	broadcast  chan Message

	mutex   sync.Locker
	clients map[string]*Client
}

func NewChatRoom(roomId string) *ChatRoom {
	return &ChatRoom{
		RoomId:     roomId,
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message, 100),
		clients:    make(map[string]*Client),
		mutex:      &sync.Mutex{},
	}
}

func (cr *ChatRoom) Run() {
	log.Println("[DEBUG] ChatRoom is running:", cr.RoomId)
	for {
		select {
		case client := <-cr.register:
			cr.mutex.Lock()
			cr.clients[client.ID] = client
			cr.mutex.Unlock()

			cr.broadcast <- Message{Sender: SystemSender, Content: client.ID + " has joined the chat"}
			log.Printf("[DEBUG] Client %s registered in room %s\n", client.ID, cr.RoomId)

		case client := <-cr.unregister:
			cr.mutex.Lock()
			delete(cr.clients, client.ID)
			cr.mutex.Unlock()

			cr.broadcast <- Message{Sender: SystemSender, Content: client.ID + " has left the chat"}
			log.Printf("[DEBUG] Client %s unregistered from room %s\n", client.ID, cr.RoomId)

		case message := <-cr.broadcast:
			for _, client := range cr.clients {
				// NOTE: broadcast message to all clients except the sender
				if message.Sender == client.ID {
					continue
				}
				if err := client.WriteMessage(fmt.Sprintf("[%s]: %s", message.Sender, message.Content)); err != nil {
					cr.unregister <- client
					client.Close()
				}
			}
		}
	}
}

func (cr *ChatRoom) RegisterClient(client *Client) {
	if client.ID == SystemSender {
		client.WriteMessage("[ERROR] Client ID cannot be " + SystemSender)
		client.Close()
		return
	}

	if _, exists := cr.clients[client.ID]; exists {
		client.WriteMessage("[ERROR] Client ID already exists")
		client.Close()
		return
	}

	cr.register <- client
}
