package main

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

func TestCreateAndDeleteUser(t *testing.T) {
	db.InitDB()

	// Step 1: Upload user with image
	filePath := filepath.Join("testdata", "sample.jpg")
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open sample image: %v", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	_ = writer.WriteField("name", "TestUser API")

	part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		t.Fatalf("Failed to copy image: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/users", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()
	handlers.UsersHandler(rr, req)

	// Step 2: Extract User ID
	body := rr.Body.String()
	start := strings.Index(body, `"ID":"`) + 6
	end := strings.Index(body[start:], `"`) + start
	userID := body[start:end]
	if userID == "" {
		t.Fatal("Failed to extract user ID")
	}

	// Step 3: Wait for PDF to be generated
	pdfPath := filepath.Join("profiles", userID+".pdf")
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(pdfPath); err == nil {
			break
		}
		time.Sleep(300 * time.Millisecond)
	}

	// Step 4: Test GET /users/{id}
	getReq := httptest.NewRequest(http.MethodGet, "/users/"+userID, nil)
	getRes := httptest.NewRecorder()
	handlers.UserHandler(getRes, getReq)
	if getRes.Code != http.StatusOK {
		t.Errorf("GET /users/{id} failed: expected 200, got %d", getRes.Code)
	}

	// Step 5: Test PUT /users/{id}
	putBody := strings.NewReader("name=UpdatedName")
	putReq := httptest.NewRequest(http.MethodPut, "/users/"+userID, putBody)
	putReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	putRes := httptest.NewRecorder()
	handlers.UserHandler(putRes, putReq)
	if putRes.Code != http.StatusOK {
		t.Errorf("PUT /users/{id} failed: expected 200, got %d", putRes.Code)
	}
	if !strings.Contains(putRes.Body.String(), `"Name":"UpdatedName"`) {
		t.Errorf("PUT failed to update user name, got: %s", putRes.Body.String())
	}

	// Step 6: Test GET /profile/{id}
	profileReq := httptest.NewRequest(http.MethodGet, "/profile/"+userID, nil)
	profileRes := httptest.NewRecorder()
	handlers.ProfileHandler(profileRes, profileReq)
	if profileRes.Code != http.StatusOK {
		t.Fatalf("PDF Download failed: got status %d", profileRes.Code)
	}
	if profileRes.Header().Get("Content-Type") != "application/pdf" {
		t.Fatalf("Expected PDF Content-Type, got %s", profileRes.Header().Get("Content-Type"))
	}

	// Step 7: Test edge cases (invalid ID)
	check404 := func(method, url string) {
		req := httptest.NewRequest(method, url, nil)
		res := httptest.NewRecorder()
		if strings.HasPrefix(url, "/profile/") {
			handlers.ProfileHandler(res, req)
		} else {
			handlers.UserHandler(res, req)
		}
		if res.Code != http.StatusNotFound {
			t.Errorf("Expected 404 for %s %s, got %d", method, url, res.Code)
		}
	}
	check404(http.MethodGet, "/users/invalid-id")
	check404(http.MethodDelete, "/users/invalid-id")
	check404(http.MethodGet, "/profile/invalid-id")

	// Step 8: Final DELETE /users/{id}
	delReq := httptest.NewRequest(http.MethodDelete, "/users/"+userID, nil)
	delRes := httptest.NewRecorder()
	handlers.UserHandler(delRes, delReq)
	if delRes.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK on final delete, got %d", delRes.Code)
	}
}
