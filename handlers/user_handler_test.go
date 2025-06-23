package handlers

import (
	"bytes"
	"encoding/json"
	"go-user-service/models"
	"go-user-service/utils"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	
)

func TestUsersHandler_GET(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()

	UsersHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
}

func TestUsersHandler_POST(t *testing.T) {
	os.MkdirAll("uploads", os.ModePerm)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("name", "Test User")
	imagePath := filepath.Join("testdata", "sample.jpg")
	file, _ := os.Open(imagePath)
	defer file.Close()

	part, _ := writer.CreateFormFile("image", filepath.Base(imagePath))
	io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest("POST", "/users", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()

	UsersHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", rr.Code)
	}

	var user models.User
	_ = json.NewDecoder(rr.Body).Decode(&user)

	if user.ID == "" || user.Name != "Test User" {
		t.Fatalf("Invalid user response: %+v", user)
	}

	// Clean up
	os.Remove(filepath.Join("uploads", filepath.Base(user.ImageURL)))
}

func TestUserHandler_CRUD(t *testing.T) {
	// Setup a dummy user
	id := utils.GenerateID()
	user := &models.User{
		ID:   id,
		Name: "Original Name",
	}
	models.Users[id] = user

	// Test GET
	getReq := httptest.NewRequest("GET", "/users/"+id, nil)
	getRes := httptest.NewRecorder()
	UserHandler(getRes, getReq)
	if getRes.Code != http.StatusOK {
		t.Errorf("GET /users/{id} failed: got %d", getRes.Code)
	}

	// Test PUT
	putData := strings.NewReader("name=Updated Name")
	putReq := httptest.NewRequest("PUT", "/users/"+id, putData)
	putReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	putRes := httptest.NewRecorder()
	UserHandler(putRes, putReq)

	if !strings.Contains(putRes.Body.String(), "Updated Name") {
		t.Errorf("PUT /users/{id} did not update name")
	}

	// Test DELETE
	delReq := httptest.NewRequest("DELETE", "/users/"+id, nil)
	delRes := httptest.NewRecorder()
	UserHandler(delRes, delReq)

	if delRes.Code != http.StatusOK {
		t.Errorf("DELETE /users/{id} failed: got %d", delRes.Code)
	}
}

func TestUserHandler_InvalidUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/invalid-id", nil)
	rr := httptest.NewRecorder()

	UserHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected 404 for invalid user ID, got %d", rr.Code)
	}
}

func TestProfileHandler(t *testing.T) {
	// Setup dummy user with PDF path
	id := utils.GenerateID()
	pdfPath := filepath.Join("profiles", id+".pdf")
	os.MkdirAll("profiles", os.ModePerm)
	os.WriteFile(pdfPath, []byte("dummy pdf content"), 0644)

	user := &models.User{ID: id, PDFPath: pdfPath}
	models.Users[id] = user

	req := httptest.NewRequest("GET", "/profile/"+id, nil)
	rr := httptest.NewRecorder()

	ProfileHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}

	if rr.Header().Get("Content-Type") != "application/pdf" {
		t.Errorf("Expected PDF Content-Type, got %s", rr.Header().Get("Content-Type"))
	}

	// Cleanup
	os.Remove(pdfPath)
}

func TestProfileHandler_InvalidUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/profile/invalid-id", nil)
	rr := httptest.NewRecorder()

	ProfileHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found for invalid profile, got %d", rr.Code)
	}
}
