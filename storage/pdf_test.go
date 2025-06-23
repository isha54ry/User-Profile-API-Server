package storage

import (
	"go-user-service/models"
	"os"
	"path/filepath"
	"testing"
)

func TestGeneratePDF(t *testing.T) {

	imagePath := filepath.Join("..","testdata", "sample.jpg")
	absPath, err := filepath.Abs(imagePath)
	if err != nil {
		t.Fatal("Failed to get absolute path:", err)
	}

	user := &models.User{
		ID:       "test123",
		Name:     "Test User",
		ImageURL: absPath, // Use absolute path for testing
		PDFPath:  "",      // Initially empty, will be set by GeneratePDF
	}

	// Act
	GeneratePDF(user)

	// Assert: check if PDF file was created
	expectedPath := filepath.Join("profiles", user.ID+".pdf")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("Expected PDF at %s, but it does not exist", expectedPath)
	}

	// Cleanup
	_ = os.Remove(expectedPath)
}
