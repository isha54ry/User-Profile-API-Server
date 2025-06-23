package handlers

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "go-user-service/db"
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

        user := &models.User{ID: id, Name: name, ImageURL: imagePath}

        _, err = db.DB.Exec(`INSERT INTO users (id, name, image_url, pdf_path) VALUES ($1, $2, $3, $4)`,
            user.ID, user.Name, user.ImageURL, "")
        if err != nil {
            http.Error(w, "Failed to save user", http.StatusInternalServerError)
            return
        }

        go storage.UploadImage(user, imagePath)
        go storage.GeneratePDF(user)

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)

    case "GET":
        rows, err := db.DB.Query(`SELECT id, name, image_url, pdf_path FROM users`)
        if err != nil {
            http.Error(w, "Database error", http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        users := make(map[string]*models.User)
        for rows.Next() {
            user := new(models.User)
            rows.Scan(&user.ID, &user.Name, &user.ImageURL, &user.PDFPath)
            users[user.ID] = user
        }
        json.NewEncoder(w).Encode(users)
    }
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/users/")

    user := new(models.User)
    err := db.DB.QueryRow(`SELECT id, name, image_url, pdf_path FROM users WHERE id = $1`, id).
        Scan(&user.ID, &user.Name, &user.ImageURL, &user.PDFPath)
    if err == sql.ErrNoRows {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(user)

    case "PUT":
        name := r.FormValue("name")
        if name != "" {
            user.Name = name
            _, err := db.DB.Exec(`UPDATE users SET name = $1 WHERE id = $2`, name, id)
            if err != nil {
                http.Error(w, "Failed to update user", http.StatusInternalServerError)
                return
            }
        }
        json.NewEncoder(w).Encode(user)

    case "DELETE":
        os.Remove(user.PDFPath)
        os.Remove(filepath.Join("uploads", filepath.Base(user.ImageURL)))
        _, err := db.DB.Exec(`DELETE FROM users WHERE id = $1`, id)
        if err != nil {
            http.Error(w, "Failed to delete user", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    }
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/profile/")

    user := new(models.User)
    err := db.DB.QueryRow(`SELECT id, name, image_url, pdf_path FROM users WHERE id = $1`, id).
        Scan(&user.ID, &user.Name, &user.ImageURL, &user.PDFPath)
    if err == sql.ErrNoRows || user.PDFPath == "" {
        http.Error(w, "Profile not ready", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    http.ServeFile(w, r, user.PDFPath)
}
