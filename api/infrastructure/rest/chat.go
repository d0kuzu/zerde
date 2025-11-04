package rest

import (
	"AISale/api/infrastructure/controllers/chat_controllers"
	"AISale/config"
	"AISale/ws"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, app *config.App) {
	h := chat_controllers.NewChatHandler(app.Cfg)
	wsh := ws.NewWSHandler(app.Cfg)

	productGroup := router.Group("chats")
	{
		productGroup.GET("/get_all", h.GetAllChats)
		// productGroup.GET("/get_chat", h.GetChat)
		productGroup.GET("/get_pagination", h.GetPagination)
		productGroup.GET("/search_chat", h.SearchChat)

		productGroup.GET("/get_conversation", wsh.ChatPolling)
	}
}
