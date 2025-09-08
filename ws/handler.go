package ws

import (
	"AISale/config"
	twilio "AISale/services/twillio"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	cfg *config.Settings
}

func NewWSHandler(cfg *config.Settings) *WSHandler {
	return &WSHandler{cfg: cfg}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *WSHandler) ChatPolling(c *gin.Context) {
	chatID := c.Query("chat")
	if chatID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chat id required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	client := &Client{conn: conn, chat: chatID}
	RegisterClient(client)
	log.Println("new client connected to chat", chatID)

	go PollTwilio(chatID, h.cfg.AccountSID, h.cfg.AuthToken)

	go client.Listen(h.cfg.AccountSID, h.cfg.AuthToken)
}

func (c *Client) Listen(accountSID, authToken string) {
	defer UnregisterClient(c)

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("client disconnected:", err)
			return
		}
		log.Printf("[Operator → Twilio] Chat %s: %s\n", c.chat, string(msg))

		twilioClient := twilio.NewClient(accountSID, authToken)
		_, err = twilioClient.SendMessage(config.BotNumber, c.chat, string(msg))
		if err != nil {
			log.Println("ws twillio message send error:", err)
		}
	}
}
