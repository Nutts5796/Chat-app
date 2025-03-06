package main

import (
	"log"
	"net/http"

	"github.com/Nutts5796/Chat-app/config"
	"github.com/Nutts5796/Chat-app/handlers"
	"github.com/Nutts5796/Chat-app/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db := repository.ConnectDB(cfg.DBUrl)
	defer db.Close()

	// Инициализация Gin
	r := gin.Default()

	// Маршруты для аутентификации
	r.POST("/register", func(c *gin.Context) {
		handlers.Register(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		handlers.Login(c, db)
	})

	// Защищенные маршруты
	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	{
		// Здесь можно добавить маршруты для чатов и сообщений
	}

	// Запуск сервера
	log.Printf("Сервер запущен на порту %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

// AuthMiddleware - middleware для проверки JWT токена
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
			c.Abort()
			return
		}

		claims := &handlers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.LoadConfig().JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
