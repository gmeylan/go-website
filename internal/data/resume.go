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
	}
}
