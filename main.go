package main

import (
	"log"
	"net/http"

    _ "github.com/lib/pq" // PostgreSQL driver
	"go-user-service/db"
	"go-user-service/handlers"
)

// Global CORS Middleware
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Respond OK to preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db.InitDB()

	// Apply middleware to all routes
	http.Handle("/users", withCORS(http.HandlerFunc(handlers.UsersHandler)))
	http.Handle("/users/", withCORS(http.HandlerFunc(handlers.UserHandler)))
	http.Handle("/profile/", withCORS(http.HandlerFunc(handlers.ProfileHandler)))

	log.Println("Server started at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
