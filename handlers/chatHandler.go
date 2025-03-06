package handlers

import (
	"chat-app/db"
	"chat-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {
	var chat models.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&chat)
	c.JSON(http.StatusOK, gin.H{"message": "Chat created successfully", "chatID": chat.ID})
}
