package handlers

import (
	"fmt"
	"time"

	"jimdel/pkg/server/helpers"
	"jimdel/pkg/web/views/pages"
	"net/http"
	"sort"
)

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/projects":
		projects := getAllProjects()
		pages.Projects(projects).Render(r.Context(), w)
	default:
		http.NotFound(w, r)
	}
}

func getAllProjects() []helpers.Project {
	files := helpers.GetAllFilesInDir("projects")
	projects := []helpers.Project{}
	for _, file := range files {
		fileName := fmt.Sprintf("projects/%s", file)
		project := helpers.ParseYaml(fileName, &helpers.Project{})
		projects = append(projects, *project)
	}
	// Sort projects by the Date field
	sort.Slice(projects, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02", projects[i].Date)
		dateJ, _ := time.Parse("2006-01-02", projects[j].Date)
		return dateJ.Before(dateI)
	})
	return projects
}
