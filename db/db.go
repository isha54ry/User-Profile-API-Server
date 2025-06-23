package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func InitDB() {
    connStr := "host=localhost port=5432 user=postgres password=22052388 dbname=go_user_service sslmode=disable"
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error opening DB:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Error connecting to DB:", err)
    }

    log.Println("Connected to PostgreSQL!")
}