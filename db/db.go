package db

import (
	"chat-app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// Инициализация базы данных
func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", config.DB_URL)
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&User{}, &Chat{}, &Message{})
}

// Закрытие соединения с базой данных
func CloseDB() {
	DB.Close()
}
