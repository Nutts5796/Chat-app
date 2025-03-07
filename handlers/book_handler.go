package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "library-app/models"
    "github.com/gorilla/mux"
)

func GetBooks(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT id, title, author, year FROM books")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var books []models.Book
        for rows.Next() {
            var b models.Book
            if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            books = append(books, b)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(books)
    }
}

func GetBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

        var b models.Book
        err = db.QueryRow("SELECT id, title, author, year FROM books WHERE id = $1", id).Scan(&b.ID, &b.Title, &b.Author, &b.Year)
        if err != nil {
            if err == sql.ErrNoRows {
                http.Error(w, "Book not found", http.StatusNotFound)
            } else {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(b)
    }
}

func CreateBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var b models.Book
        if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        err := db.QueryRow("INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id", b.Title, b.Author, b.Year).Scan(&b.ID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(b)
    }
}

func UpdateBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

        var b models.Book
        if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        _, err = db.Exec("UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4", b.Title, b.Author, b.Year, id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func DeleteBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

        _, err = db.Exec("DELETE FROM books WHERE id = $1", id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}