package chat_controllers

import (
	"AISale/config"
	"AISale/services/airtable"
	twilio "AISale/services/twillio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ChatHandler struct {
	cfg *config.Settings
}

func NewChatHandler(cfg *config.Settings) *ChatHandler {
	return &ChatHandler{cfg: cfg}
}

func (h *ChatHandler) GetAllChats(c *gin.Context) {
	client := airtable.NewClient(h.cfg.ApiKey, h.cfg.BaseID)

	records, err := client.ListRecords(h.cfg.TableName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"answer": string(records)})
}

func (h *ChatHandler) GetChat(c *gin.Context) {
	clientNumber := c.Query("chat")
	if clientNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client number required"})
		return
	}
	botNumber := "+16693420294"

	twilioClient := twilio.NewClient(h.cfg.AccountSID, h.cfg.AuthToken)

	messages, err := twilioClient.GetConversation(clientNumber, botNumber, 50)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range messages {
		fmt.Printf("[%s] %s → %s: %s\n", m.DateCreated, m.From, m.To, m.Body)
	}
}
