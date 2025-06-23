package handlers
// handlers/user_handlers.go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"go-user-service/models"
	"go-user-service/storage"
	"go-user-service/utils"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		name := r.FormValue("name")
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Image required", http.StatusBadRequest)
			return
		}
		defer file.Close()

		id := utils.GenerateID()
		imagePath := fmt.Sprintf("uploads/%s_%s", id, header.Filename)
		os.MkdirAll("uploads", os.ModePerm)
		out, _ := os.Create(imagePath)
		io.Copy(out, file)

		user := &models.User{ID: id, Name: name}
		models.Users[id] = user

		go storage.UploadImage(user, imagePath)
		go storage.GeneratePDF(user)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	case "GET":
		json.NewEncoder(w).Encode(models.Users)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	user, ok := models.Users[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(user)
	case "PUT":
		name := r.FormValue("name")
		if name != "" {
			user.Name = name
		}
		models.Users[id] = user
		json.NewEncoder(w).Encode(user)
	case "DELETE":
		delete(models.Users, id)
		os.Remove(user.PDFPath)
		os.Remove(filepath.Join("uploads", filepath.Base(user.ImageURL)))
		w.WriteHeader(http.StatusOK)
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/profile/")
	user, ok := models.Users[id]
	if !ok || user.PDFPath == "" {
		http.Error(w, "Profile not ready", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, user.PDFPath)
}
