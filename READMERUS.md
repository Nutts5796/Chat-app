![Скриншот](images/Golang.png)
# Library App
Простое REST API для управления книгами в библиотеке, написанное на Go. Проект использует Docker для контейнеризации, PostgreSQL для хранения данных и Makefile для автоматизации команд.

# 📋 О проекте

Этот проект представляет собой бэкенд-приложение для управления книгами в библиотеке. Оно предоставляет REST API для выполнения CRUD-операций (создание, чтение, обновление, удаление) с книгами.

Основные функции:

1) Получение списка всех книг.
2) Получение информации о конкретной книге по её ID.
3) Добавление новой книги.
4) Обновление информации о существующей книге.
5) Удаление книги.

# 🛠  Технологии

Go — язык программирования для бэкенда.
Docker — для контейнеризации приложения и базы данных.
PostgreSQL — реляционная база данных для хранения информации о книгах.
Makefile — для автоматизации команд сборки и запуска.

# 🚀 Быстрый старт

Предварительные требования

Установите Docker.
Установите Docker Compose.
Убедитесь, что у вас установлен Go (опционально, если вы хотите запускать проект локально без Docker).

# Установка и запуск

1) Клонируйте репозиторий:
   git clone https://github.com/Nutts5796/library-app.git
   cd library-app
2) Соберите и запустите проект с помощью Docker:
    make build
    make run
3) Примените миграции для создания таблицы в базе данных:
    make migrate
4) Приложение будет доступно по адресу: http://localhost:8080.

# 📚 API Endpoints

Получить список всех книг

Метод: GET
URL: /books
Пример ответа:
    [
  {
    "id": 1,
    "title": "1984",
    "author": "George Orwell",
    "year": 1949
  },
  {
    "id": 2,
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "year": 1960
  }
]

# Получить информацию о книге по ID
Метод: GET
URL: /books/{id}
Пример ответа:
    {
  "id": 1,
  "title": "1984",
  "author": "George Orwell",
  "year": 1949
}

# Добавить новую книгу

Метод: POST
URL: /books
Тело запроса:

{
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "year": 1925
}

Пример ответа:
    {
  "id": 3,
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "year": 1925
}

# Обновить информацию о книге

Метод: PUT
URL: /books/{id}
Тело запроса:
    {
  "title": "1984",
  "author": "George Orwell",
  "year": 1948
}

# Удалить книгу

Метод: DELETE
URL: /books/{id}

# 🐳 Docker

Проект использует Docker для контейнеризации. В репозитории есть Dockerfile и docker-compose.yml, которые позволяют легко запустить приложение и базу данных.

Команды Docker

Собрать образ:
    docker-compose build
Запустить контейнеры:
    docker-compose up
Остановить контейнеры:
    docker-compose down

# 🛠 Makefile

Для удобства в проекте используется Makefile. Вот основные команды:
Собрать образ:
    make build
Запустить контейнеры:
    make run
Остановить контейнеры:
    make down
Остановить контейнеры:
    make down

# 📁 Структура проекта

library-app/
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
├── main.go
├── models/
│   └── book.go
├── handlers/
│   └── book_handler.go
├── migrations/
│   └── 001_create_books_table.sql
└── docker-compose.yml
