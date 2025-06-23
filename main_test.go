package main

import (
	"net/http"
	"testing"
	"time"
)

func TestMainServerStartup(t *testing.T) {
	go func() {
		StartServer(":8082") // Use a different port to avoid conflict
	}()

	// Give the server a moment to start
	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get("http://localhost:8082/users")
	if err != nil {
		t.Fatalf("Failed to reach server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Unexpected response code: got %d", resp.StatusCode)
	}
}
