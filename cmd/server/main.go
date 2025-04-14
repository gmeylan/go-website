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

	router.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		h.About(w, r)
	})

	router.HandleFunc("GET /portfolio", func(w http.ResponseWriter, r *http.Request) {
		h.Portfolio(w, r)
	})

	router.HandleFunc("GET /technologies", func(w http.ResponseWriter, r *http.Request) {
		h.Technologies(w, r)
	})

	router.HandleFunc("GET /blog", func(w http.ResponseWriter, r *http.Request) {
		h.Blog(w, r)
	})

	router.HandleFunc("GET /blog/{slug}", func(w http.ResponseWriter, r *http.Request) {
		h.BlogPost(w, r)
	})

	router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info("Starting server on :8080")
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("Server error", slog.Any("error", err))
		os.Exit(1)
	}

}
