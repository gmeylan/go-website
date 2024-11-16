package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gmeylan/go-website/internal/database"
)

func main() {
	dbPath := filepath.Join(".", "test.db")
	db, err := database.NewSQLiteDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Printf("Successfully connected to database at %s", dbPath)

	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	log.Println("Server starting on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
