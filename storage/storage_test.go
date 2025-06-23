package storage

import (
	"go-user-service/models"
	"os"
	"path/filepath"
	"testing"
)

func TestGeneratePDFStorage(t *testing.T) {
	imagePath := filepath.Join("..", "testdata", "sample.jpg")
	absPath, err := filepath.Abs(imagePath)
	if err != nil {
		t.Fatal("Failed to get absolute path:", err)
	}

	user := &models.User{
		ID:       "test123",
		Name:     "Test User",
		ImageURL: absPath,
	}

	GeneratePDF(user)

	pdfPath := filepath.Join("profiles", user.ID+".pdf")
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		t.Errorf("PDF not generated at %s", pdfPath)
	}

	// Clean up
	_ = os.Remove(pdfPath)
}
