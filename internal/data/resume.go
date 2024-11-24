package data

import (
	"github.com/gmeylan/go-website/internal/types"
)

func GetExperiences() []types.Experience {
	return []types.Experience{
		{
			Title:       "Software Engineer",
			Company:     "EPFL",
			Location:    "Lausanne",
			Period:      "Janvier 2021 - Présent",
			Description: "Développement et optimisation de solutions web pour la gestion des candidatures doctorales et des ressources informatiques.",
			Tags:        []string{"Python", "Django", "Docker", "Kubernetes", "Bash", "..."},
		},
		{
			Title:       "Ingénieur devOps",
			Company:     "tiko Energy Solutions",
			Location:    "Zurich",
			Period:      "Mai 2019 - Décembre 2020",
			Description: "Conception et implémentation de solutions pour l'aide aux équipes operationnelles pour une solutions d'aide dans la gestion d'énergie intelligente.",
			Tags:        []string{"Python", "Django", "Docker", "AWS", "Bash", "..."},
		},
		{
			Title:       "Développeur Full Stack",
			Company:     "Propulse Lab",
			Location:    "Lausanne",
			Period:      "Juillet 2018 - Avril 2019",
			Description: "Conception et implémentation de solutions pour le domaine du marketing digital direct.",
			Tags:        []string{"Python", "Django", "Docker", "AWS", "Bash", "JS/TS"},
		},
		{
			Title:       "Développeur Python",
			Company:     "Open Net",
			Location:    "Lausanne",
			Period:      "Septembre 2017 - Juin 2018",
			Description: "Conception et implémentation de solutions dans le framework ERP Odoo.",
			Tags:        []string{"Python", "Odoo", "Docker", "Bash"},
		},
		{
			Title:       "Développeur Full Stack",
			Company:     "Divers entreprises",
			Location:    "Lausanne",
			Period:      "Septembre 2016 - Juin 2016",
			Description: "Conception et implémentation de solutions dans des frameworks PHP.",
			Tags:        []string{"PHP", "Symfony", "Laravel", "JS/TS"},
		},
	}
}
