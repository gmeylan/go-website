package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gmeylan/go-website/internal/components"
	"github.com/gmeylan/go-website/internal/components/about"
	"github.com/gmeylan/go-website/internal/components/blog"
	"github.com/gmeylan/go-website/internal/components/portfolio"
	"github.com/gmeylan/go-website/internal/components/technologies"
	"github.com/gmeylan/go-website/internal/data"
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
	experiences := data.GetExperiences()
	component := about.About(experiences)
	component.Render(r.Context(), w)
}

func (h *Handlers) Blog(w http.ResponseWriter, r *http.Request) {
	blog.Blog().Render(r.Context(), w)
}

func (h *Handlers) BlogPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	blog.BlogPost(slug).Render(r.Context(), w)
}

func (h *Handlers) Portfolio(w http.ResponseWriter, r *http.Request) {

	portfolio.Portfolio().Render(r.Context(), w)
}

func (h *Handlers) Technologies(w http.ResponseWriter, r *http.Request) {
	technologies.Technologies().Render(r.Context(), w)
}
