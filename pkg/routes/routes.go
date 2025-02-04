package routes

import (
	"my-go-api/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/receipts/process", handlers.PostPoints)
	router.GET("/receipts/:id/points", handlers.GetID)
}
