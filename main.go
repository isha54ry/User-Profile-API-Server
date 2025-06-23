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

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ✅ Extracted StartServer so it can be tested
func StartServer(port string) {
	db.InitDB()

	mux := http.NewServeMux()
	mux.Handle("/users", withCORS(http.HandlerFunc(handlers.UsersHandler)))
	mux.Handle("/users/", withCORS(http.HandlerFunc(handlers.UserHandler)))
	mux.Handle("/profile/", withCORS(http.HandlerFunc(handlers.ProfileHandler)))

	log.Printf("Server started at http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func main() {
	StartServer(":8081")
}
