package rest

import (
	"AISale/api/infrastructure/controllers/chat_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, settings *config.Settings) {
	h := chat_controllers.NewChatHandler(settings)

	productGroup := router.Group("chats")
	{
		productGroup.GET("/get_all", h.GetAllChats)
	}
}
