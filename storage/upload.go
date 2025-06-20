package storage

import (
    "fmt"
    "go-user-service/models"
)

func UploadImage(user *models.User, path string) {
    // Simulate upload: just save to local and set ImageURL
    user.ImageURL = fmt.Sprintf("/%s", path)
}