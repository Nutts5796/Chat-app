package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(dataSourceName string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	return db
}
