package chat_controllers

import (
	"AISale/config"
	twilio "AISale/services/twillio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type ChatHandler struct {
	cfg *config.Settings
}

func NewChatHandler(cfg *config.Settings) *ChatHandler {
	return &ChatHandler{cfg: cfg}
}

func (h *ChatHandler) GetAllChats(c *gin.Context) {
	// get numbers from airtable

	//client := airtable.NewClient(h.cfg.ApiKey, h.cfg.BaseID)
	//
	//records, err := client.ListRecords(h.cfg.TableName)
	//if err != nil {
	//	c.JSON(500, gin.H{"error": err.Error()})
	//	return
	//}

	// get messages from twillio

	clientNumber := "+14086851938"
	botNumber := "+16693420294"

	twilioClient := twilio.NewClient(h.cfg.AccountSID, h.cfg.AuthToken)

	messages, err := twilioClient.GetConversation(clientNumber, botNumber, 50)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range messages {
		fmt.Printf("[%s] %s → %s: %s\n", m.DateCreated, m.From, m.To, m.Body)
	}

	c.JSON(200, gin.H{"answer": ""})
}
