package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gmeylan/go-website/internal/database"
	"github.com/gmeylan/go-website/internal/handlers"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbPath := filepath.Join(".", "test.db")
	db, err := database.NewSQLiteDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	logger.Info("Successfully connected to database")

	router := http.NewServeMux()

	h := handlers.NewHandlers(db)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		h.Home(w, r)
	})

	router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info("Server listening on :8080")

	srv.ListenAndServe()
}
