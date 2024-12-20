package server

import (
	"fmt"
	pages "jimdel/pkg/web/views/pages"

	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Route struct {
	Title     string
	Component templ.Component
	Path      string
}

type Routes []Route

// Register pages
var routes = Routes{
	Route{
		Title:     "Home",
		Component: pages.Home(),
		Path:      "/",
	},
	Route{
		Title:     "About",
		Component: pages.About(),
		Path:      "/about",
	},
	Route{
		Title:     "Blog",
		Component: pages.Blog(),
		Path:      "/blog",
	},
	Route{
		Title:     "Contact",
		Component: pages.Contact(),
		Path:      "/contact",
	},
}

func Run(PORT string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	for _, route := range routes {
		r.Get(route.Path, func(w http.ResponseWriter, r *http.Request) {
			route.Component.Render(r.Context(), w)
		})
	}

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		component := pages.NotFound()
		component.Render(r.Context(), w)
	})

	fmt.Printf("Server is running on port: %s", PORT)
	err := http.ListenAndServe(PORT, r)
	fmt.Printf("Server error %v", err)

	return err
}
