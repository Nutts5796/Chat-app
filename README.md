![Скриншот](images/Go.png)

# Library App
A simple REST API for managing books in a library, written in Go. The project uses Docker for containerization, PostgreSQL for data storage, and Makefile for command automation.

# 📋 About the Project

This project is a backend application for managing books in a library. It provides a REST API for performing CRUD operations (create, read, update, delete) on books.

Key Features:

Retrieve a list of all books.
Get details of a specific book by its ID.
Add a new book.
Update information about an existing book.
Delete a book.

# 🛠 Technologies

Go — programming language for the backend.
Docker — for containerizing the application and database.
PostgreSQL — relational database for storing book information.
Makefile — for automating build and run commands.

# 🚀 Quick Start

Prerequisites

Install Docker.
Install Docker Compose.
Ensure you have Go installed (optional, if you want to run the project locally without Docker).
Installation and Setup

Clone the repository:
    git clone https://github.com/Nutts5796/library-app.git
    cd library-app
Build and run the project using Docker:
    make build
    make run
Apply migrations to create the table in the database:
    make migration
The application will be available at: http://localhost:8080.

# 🐳 Docker

The project uses Docker for containerization. The repository includes a Dockerfile and docker-compose.yml to easily run the application and database.

Docker Commands

Build the image:
    docker-compose build
Run the containers:
    docker-compose up
Stop the containers:
    docker-compose down

# 🛠 Makefile

For convenience, the project uses a Makefile. Here are the main commands:

Build the project:
    make build
Run the project:
    make run
Stop the containers:
    make down
Apply migrations:
    make migrate
│   └── 001_create_books_table.sql
└── docker-compose.yml
