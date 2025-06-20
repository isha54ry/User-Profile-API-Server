package storage

import (
    "fmt"
    "go-user-service/models"
    "github.com/jung-kurt/gofpdf"
    "path/filepath"
)

func GeneratePDF(user *models.User) {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, fmt.Sprintf("Name: %s", user.Name))
    pdf.Ln(20)
    pdf.Image(user.ImageURL[1:], 10, 30, 50, 0, false, "", 0, "")

    path := filepath.Join("profiles", fmt.Sprintf("%s.pdf", user.ID))
    pdf.OutputFileAndClose(path)
    user.PDFPath = path
}