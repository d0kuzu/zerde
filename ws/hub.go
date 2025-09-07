package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	chat string
}

var (
	clients   = make(map[*Client]bool)
	clientsMu sync.Mutex
)

func RegisterClient(c *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[c] = true
}

func UnregisterClient(c *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, c)
	c.conn.Close()
}

func Broadcast(chatID string, msg []byte) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for c := range clients {
		if c.chat == chatID {
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("ws message write error:", err)
				c.conn.Close()
				delete(clients, c)
			}
		}
	}
}
