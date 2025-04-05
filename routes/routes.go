package routes

import (
	"github.com/vknow360/FcmGo/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/send/topic/notification", handlers.SendNotificationToTopic)
	r.POST("/send/topic/message", handlers.SendMessageToTopic)
	r.POST("/send/tokens/notification", handlers.SendNotificationToTokens)
	r.POST("/send/tokens/message", handlers.SendMessageToTokens)
	r.GET("/status", handlers.Status)
	return r
}
