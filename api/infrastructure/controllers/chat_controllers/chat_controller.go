package chat_controllers

import (
	"AISale/config"
	"AISale/services/airtable"
	twilio "AISale/services/twillio"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type ChatHandler struct {
	cfg *config.Settings
}

func NewChatHandler(cfg *config.Settings) *ChatHandler {
	return &ChatHandler{cfg: cfg}
}

func (h *ChatHandler) GetAllChats(c *gin.Context) {
	client := airtable.NewClient(h.cfg.ApiKey, h.cfg.BaseID)

	page, err := strconv.ParseInt(c.Query("page"), 10, 32)
	if err != nil {
		panic(err)
	}

	records, err := client.ListPageRecords(h.cfg.TableName, page, 20)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"answer": records})
}

func (h *ChatHandler) GetPagination(c *gin.Context) {
	client := airtable.NewClient(h.cfg.ApiKey, h.cfg.BaseID)

	pages, err := client.GetTotalPages(h.cfg.TableName, 20)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"answer": pages})
}

func (h *ChatHandler) GetChat(c *gin.Context) {
	clientNumber := c.Query("chat")

	twilioClient := twilio.NewClient(h.cfg.AccountSID, h.cfg.AuthToken)

	messages, err := twilioClient.GetConversation(clientNumber, config.BotNumber, 1000)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"answer": messages})
}
