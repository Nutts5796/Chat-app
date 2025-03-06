package models

type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
}

type Chat struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Message struct {
	ID      int    `db:"id" json:"id"`
	ChatID  int    `db:"chat_id" json:"chat_id"`
	UserID  int    `db:"user_id" json:"user_id"`
	Content string `db:"content" json:"content"`
}
