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

func (h *UserHandler) ChangePrompt(c *gin.Context) {
	err := h.chrome.ChangePrompt(h.cfg.DiaxelLogin, h.cfg.DiaxelPassword, "test")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}
