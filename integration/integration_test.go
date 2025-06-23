package integration

import (
	"bytes"
	"go-user-service/db"
	"go-user-service/handlers"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func setupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handlers.UsersHandler)
	mux.HandleFunc("/users/", handlers.UserHandler)
	mux.HandleFunc("/profile/", handlers.ProfileHandler)
	return mux
}

func TestUserLifecycleIntegration(t *testing.T) {
	db.InitDB()

	server := httptest.NewServer(setupRouter())
	defer server.Close()

	// Prepare multipart form
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	_ = writer.WriteField("name", "IntegrationTest User")

	filePath := filepath.Join("..", "testdata", "sample.jpg")
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Sample image not found: %v", err)
	}
	defer file.Close()

	part, _ := writer.CreateFormFile("image", "sample.jpg")
	io.Copy(part, file)
	writer.Close()

	// POST /users
	resp, err := http.Post(server.URL+"/users", writer.FormDataContentType(), &buf)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", resp.StatusCode)
	}

	// Read user ID from response
	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)
	start := strings.Index(body, `"ID":"`) + 6
	end := strings.Index(body[start:], `"`) + start
	userID := body[start:end]

	// Wait for PDF generation
	pdfPath := filepath.Join("profiles", userID+".pdf")
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(pdfPath); err == nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	// GET /profile/{id}
	profileResp, err := http.Get(server.URL + "/profile/" + userID)
	if err != nil {
		t.Fatalf("Failed to fetch profile: %v", err)
	}
	defer profileResp.Body.Close()

	if profileResp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK for profile, got %d", profileResp.StatusCode)
	}
	if profileResp.Header.Get("Content-Type") != "application/pdf" {
		t.Errorf("Expected PDF response, got %s", profileResp.Header.Get("Content-Type"))
	}
}
