package server

import (
	"fmt"
	"io/fs"
	"jimdel/pkg/server/handlers"
	"jimdel/pkg/server/helpers"
	"jimdel/pkg/web"
	pages "jimdel/pkg/web/views/pages"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run(PORT string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))
	// Get the sub-filesystem for the static directory
	staticFS, err := fs.Sub(web.StaticFiles, "static")
	if err != nil {
		panic(err)
	}

	// Mount static file server
	helpers.FileServer(r, "/static", staticFS)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := pages.Home()
		component.Render(r.Context(), w)
	})

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		component := pages.About()
		component.Render(r.Context(), w)
	})

	r.Get("/projects", func(w http.ResponseWriter, r *http.Request) {
		handlers.ProjectHandler(w, r)
	})

	// r.Get("/blog", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.BlogHandler(w, r)
	// })

	// r.Get("/blog/{articleId}", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.BlogHandler(w, r)
	// })

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		component := pages.NotFound()
		component.Render(r.Context(), w)
	})

	fmt.Printf("Server is running on port: %s", PORT)
	err = http.ListenAndServe(PORT, r)
	fmt.Printf("Server error %v", err)

	return err
}
