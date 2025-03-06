package main

import (
	"chat-app/db"
	"chat-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/chat/:chatID", handlers.GetChatMessages)
	r.GET("/ws/:chatID", handlers.WebSocketHandler)

	r.Run(":8080")
}
