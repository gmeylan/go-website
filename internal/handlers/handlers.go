package handlers

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gmeylan/go-website/internal/components"
	"github.com/gmeylan/go-website/internal/components/about"
	"github.com/gmeylan/go-website/internal/components/blog"
	"github.com/gmeylan/go-website/internal/components/portfolio"
	"github.com/gmeylan/go-website/internal/components/technologies"
	"github.com/gmeylan/go-website/internal/data"
	"github.com/gmeylan/go-website/internal/markdown"
)

type Handlers struct {
	db     *sql.DB
	Logger *slog.Logger
}

func NewHandlers(db *sql.DB, logger *slog.Logger) *Handlers {
	return &Handlers{
		db:     db,
		Logger: logger,
	}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	posts, err := markdown.GetAllBlogPosts("content/blog/posts/", "", h.Logger)
	if err != nil {
		h.Logger.Error(err.Error())
	}

	component := components.Home(posts)
	component.Render(r.Context(), w)
}

func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	experiences := data.GetExperiences()
	component := about.About(experiences)
	component.Render(r.Context(), w)
}

func (h *Handlers) Blog(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info(r.URL.Query().Get("tag"))

	tag := r.URL.Query().Get("tag")

	posts, err := markdown.GetAllBlogPosts("content/blog/posts/", tag, h.Logger)
	if err != nil {
		h.Logger.Error(err.Error())
	}

	tags := markdown.GetAllTags(posts)

	component := blog.Blog(posts, tags)

	component.Render(r.Context(), w)
}

func (h *Handlers) BlogPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, err := markdown.ParseBlogPost(fmt.Sprintf("content/blog/posts/%s.md", slug))
	h.Logger.Info(fmt.Sprintf("content/blog/posts/%s.md", slug))

	if err != nil {
		h.Logger.Error(err.Error())
	}

	blog.BlogPost(post).Render(r.Context(), w)
}

func (h *Handlers) Portfolio(w http.ResponseWriter, r *http.Request) {
	projects := data.GetProjects()
	component := portfolio.Portfolio(projects)
	component.Render(r.Context(), w)
}

func (h *Handlers) Technologies(w http.ResponseWriter, r *http.Request) {
	technologies.Technologies().Render(r.Context(), w)
}
