package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WebSocket обработчик
func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		// Обрабатываем сообщения от клиента
		conn.WriteMessage(messageType, p)
	}
}
