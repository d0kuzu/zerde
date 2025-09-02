package rest

import (
	"AISale/api/infrastructure/controllers/chat_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, settings config.Settings) {
	productGroup := router.Group("chat")
	{
		productGroup.GET("/test", chat_controllers.Test)
	}
}
