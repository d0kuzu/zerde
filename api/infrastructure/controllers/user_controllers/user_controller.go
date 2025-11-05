package user_controllers

import (
	"AISale/config"
	"AISale/services/chrome"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	chrome *chrome.Client
	cfg    *config.Settings
}

func NewUserHandler(cfg *config.Settings, chrome *chrome.Client) *UserHandler {
	return &UserHandler{
		cfg:    cfg,
		chrome: chrome,
	}
}

type UploadTextRequest struct {
	Content string `json:"content"`
}

func (h *UserHandler) ChangePrompt(c *gin.Context) {
	var req UploadTextRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.chrome.ChangePrompt(req.Content)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

func (h *UserHandler) GetPrompt(c *gin.Context) {
	prompt, err := h.chrome.GetPrompt()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"answer": prompt})
}
