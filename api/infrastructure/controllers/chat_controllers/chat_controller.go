package chat_controllers

import (
	"AISale/api/infrastructure/response_models"
	"AISale/config"
	"AISale/services/airtable"
	twilio "AISale/services/twillio"
	"github.com/gin-gonic/gin"
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
		c.JSON(400, gin.H{"error": "invalid page parameter"})
		return
	}

	records, err := client.ListPageRecords(h.cfg.TableName, page, 10)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch Airtable records", "details": err.Error()})
		return
	}

	twilioClient := twilio.NewClient(h.cfg.AccountSID, h.cfg.AuthToken)
	twilioClient.AddMessagesCounters(records)

	var chats []response_models.ResponseRecord
	for _, record := range records {
		chats = append(chats, record.ToResponse())
	}

	c.JSON(200, gin.H{"answer": chats})
}

func (h *ChatHandler) GetPagination(c *gin.Context) {
	client := airtable.NewClient(h.cfg.ApiKey, h.cfg.BaseID)

	pages, err := client.GetTotalPages(h.cfg.TableName, 10)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"answer": pages})
}

//func (h *ChatHandler) GetChat(c *gin.Context) {
//	clientNumber := c.Query("chat")
//
//	twilioClient := twilio.NewClient(h.cfg.AccountSID, h.cfg.AuthToken)
//
//	messages, err := twilioClient.GetConversation(clientNumber, config.BotNumber, 1000)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	c.JSON(200, gin.H{"answer": messages})
//}

func (h *ChatHandler) SearchChat(c *gin.Context) {
	client := airtable.NewClient(h.cfg.ApiKey, h.cfg.BaseID)

	chat := c.Query("chat")

	records, err := client.SearchChat(h.cfg.TableName, chat)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	twilioClient := twilio.NewClient(h.cfg.AccountSID, h.cfg.AuthToken)
	twilioClient.AddMessagesCounters(records)

	var chats []response_models.ResponseRecord
	for _, record := range records {
		chats = append(chats, record.ToResponse())
	}

	c.JSON(200, gin.H{"answer": chats})
}
