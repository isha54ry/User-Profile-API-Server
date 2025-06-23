package storage

import (
	"fmt"
	"go-user-service/models"
)

func UploadImage(user *models.User, path string) {
	// Simulate upload: just save to local and set ImageURL
	user.ImageURL = path 
	fmt.Println("Image uploaded for user:", user.ID, "at", user.ImageURL)
}
