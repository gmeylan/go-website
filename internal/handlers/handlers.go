package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gmeylan/go-website/internal/components"
	"github.com/gmeylan/go-website/internal/components/about"
	"github.com/gmeylan/go-website/internal/components/portfolio"
	"github.com/gmeylan/go-website/internal/components/technologies"
)

type Handlers struct {
	db *sql.DB
}

func NewHandlers(db *sql.DB) *Handlers {
	return &Handlers{db: db}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	components.Home().Render(r.Context(), w)
}

func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	about.About().Render(r.Context(), w)
}

func (h *Handlers) Blog(w http.ResponseWriter, r *http.Request) {
	components.Blog().Render(r.Context(), w)
}

func (h *Handlers) BlogPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	components.BlogPost(slug).Render(r.Context(), w)
}

func (h *Handlers) Portfolio(w http.ResponseWriter, r *http.Request) {

	portfolio.Portfolio().Render(r.Context(), w)
}

func (h *Handlers) Technologies(w http.ResponseWriter, r *http.Request) {
	technologies.Technologies().Render(r.Context(), w)
}
