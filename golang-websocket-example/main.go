package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	upgrader := websocket.Upgrader{
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
			http.Error(w, reason.Error(), status)
		},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: true,
	}

	chatRoom := NewChatRoom("default-room")
	go chatRoom.Run()

	http.HandleFunc("/", ServeHome)
	http.HandleFunc("/ws", ServeWs(upgrader, chatRoom))

	log.Println("WebSocket server started on :8080")
	_ = http.ListenAndServe(":8080", nil)
}

func ServeWs(upgrader websocket.Upgrader, chatRoom *ChatRoom) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
			return
		}

		clientID := r.URL.Query().Get("clientId")
		if clientID == "" {
			http.Error(w, "Missing clientId", http.StatusBadRequest)
			return
		}

		client := NewClient(clientID, conn)
		client.Join(chatRoom)

		go client.readLoop()
	}
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}
