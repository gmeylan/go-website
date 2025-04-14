package data

import (
	"github.com/gmeylan/go-website/internal/types"
)

func GetProjects() []types.Portfolio {
	return []types.Portfolio{
		{
			Title:               "Mammouth",
			Tags:                []string{"Python", "Django", "Docker", "HTMX", "Alpine.js", "Tailwind"},
			Description:         "Interface de gestion du portail HaaS",
			Professional:        true,
			Website:             []string{},
			IsWebsiteAccessible: false,
			Company:             "EPFL",
		},
		{
			Title:               "EtuDB",
			Tags:                []string{"Python", "Django", "Docker", "HTMX", "Alpine.js", "Celery", "Bootstrap"},
			Description:         "Portail de gestion des candidatures de l'école doctorale IC",
			Professional:        true,
			Website:             []string{},
			IsWebsiteAccessible: false,
			Company:             "EPFL",
		},
		{
			Title:               "Go-Website",
			Tags:                []string{"Go", "Air", "Templ", "HTMX", "Tailwind"},
			Description:         "This website",
			Professional:        false,
			Website:             []string{"github.com/gmeylan/go-website", "https://meylan.io"},
			IsWebsiteAccessible: true,
			Company:             "Projet personnel",
		},
	}
}
