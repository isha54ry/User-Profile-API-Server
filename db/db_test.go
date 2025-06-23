package db

import (
    "testing"
)

func TestInitDB(t *testing.T) {
    InitDB()

    if DB == nil {
        t.Fatal("Expected DB to be initialized, got nil")
    }

    if err := DB.Ping(); err != nil {
        t.Fatalf("Expected successful ping to DB, got error: %v", err)
    }
}