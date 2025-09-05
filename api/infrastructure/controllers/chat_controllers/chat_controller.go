package chat_controllers

import (
	"AISale/config"
	"AISale/services/airtable"
	"fmt"
	"github.com/gin-gonic/gin"
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

	for _, rec := range records {
		fmt.Println(rec.ID, *rec.Fields.Email, *rec.Fields.FullName)
	}

	c.JSON(200, gin.H{"answer": records})
}
