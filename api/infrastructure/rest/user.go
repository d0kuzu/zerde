package rest

import (
	"AISale/api/infrastructure/controllers/user_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, app *config.App) {
	h := user_controllers.NewUserHandler(app.Cfg, app.Chrome)

	productGroup := router.Group("user")
	{
		productGroup.POST("/change_prompt", h.ChangePrompt)
		productGroup.GET("/get_prompt", h.GetPrompt)
	}
}
