package chat_controllers

import (
	"AISale/api/infrastructure/response_models"
	"AISale/config"
	"AISale/services/airtable"
	twilio "AISale/services/twillio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"sync"
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

	var wg sync.WaitGroup
	sem := make(chan struct{}, 5)
	errs := make(chan error, len(records))

	for i := range records {
		wg.Add(1)
		sem <- struct{}{}

		go func(rec *airtable.Record) {
			defer wg.Done()
			defer func() { <-sem }()

			messagesCounter, err := twilioClient.GetMessagesCounter(rec.Fields.MobileNumber, config.BotNumber, 1000)
			if err != nil {
				errs <- fmt.Errorf("twilio error for %s: %w", rec.Fields.MobileNumber, err)
				return
			}

			rec.Fields.MessagesCounter = messagesCounter
		}(&records[i])
	}

	wg.Wait()
	close(errs)

	var failed []string
	for err := range errs {
		log.Println(err)
		failed = append(failed, err.Error())
	}

	var chats []response_models.ResponseRecord
	for _, record := range records {
		chats = append(chats, record.ToResponse())
	}

	response := gin.H{"answer": chats}
	if len(failed) > 0 {
		response["errors"] = failed
	}

	c.JSON(200, response)
}

func (h *ChatHandler) GetPagination(c *gin.Context) {
	client := airtable.NewClient(h.cfg.ApiKey, h.cfg.BaseID)

	pages, err := client.GetTotalPages(h.cfg.TableName, 10)
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

func (h *ChatHandler) SearchChat(c *gin.Context) {
	client := airtable.NewClient(h.cfg.ApiKey, h.cfg.BaseID)

	chat := c.Query("chat")

	records, err := client.SearchChat(h.cfg.TableName, chat)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"answer": records})
}
