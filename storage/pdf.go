package storage

import (
	"fmt"
	"go-user-service/models"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(user *models.User) {
	// Ensure output directory exists
	os.MkdirAll("profiles", os.ModePerm)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %s", user.Name))
	pdf.Ln(20)
	// Use image only if valid
	if user.ImageURL != "" {
		pdf.Image(user.ImageURL, 10, 30, 50, 0, false, "", 0, "")
	}

	path := filepath.Join("profiles", fmt.Sprintf("%s.pdf", user.ID))
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		fmt.Println("Error generating PDF:", err)
	} else {
		fmt.Println("PDF saved to:", path)
	}

	user.PDFPath = path
}
