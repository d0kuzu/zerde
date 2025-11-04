package api

import (
	"AISale/api/infrastructure/rest"
	"AISale/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func RouterStart(app *config.App) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:       12 * 60 * 60,
	}))

	rest.ChatRoutes(r, app)
	rest.UserRoutes(r, app)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
