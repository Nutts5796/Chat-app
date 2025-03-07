package main

import (
    "database/sql"
    "log"
    "net/http"
    "library-app/handlers"
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

func main() {
    // Подключение к базе данных
    connStr := "user=postgres dbname=library sslmode=disable password=postgres host=db"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Инициализация маршрутизатора
    r := mux.NewRouter()

    // Регистрация обработчиков
    r.HandleFunc("/books", handlers.GetBooks(db)).Methods("GET")
    r.HandleFunc("/books/{id}", handlers.GetBook(db)).Methods("GET")
    r.HandleFunc("/books", handlers.CreateBook(db)).Methods("POST")
    r.HandleFunc("/books/{id}", handlers.UpdateBook(db)).Methods("PUT")
    r.HandleFunc("/books/{id}", handlers.DeleteBook(db)).Methods("DELETE")

    // Запуск сервера
    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}