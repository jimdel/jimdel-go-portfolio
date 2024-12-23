package handlers

import (
	"fmt"

	"jimdel/pkg/server/helpers"
	"jimdel/pkg/web/views/pages"

	"net/http"
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
	files := helpers.GetAllFilesInDir("/content/projects")
	projects := []helpers.Project{}
	for _, file := range files {
		fileName := fmt.Sprintf("/content/projects/%s", file)
		project := helpers.ParseYaml(fileName, &helpers.Project{})
		projects = append(projects, *project)
	}
	fmt.Println(projects)
	return projects
}
