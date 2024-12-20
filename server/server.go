package server

import (
	"fmt"
	pages "jimdel/pkg/web/views/pages"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run(PORT string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := pages.Home()
		component.Render(r.Context(), w)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		component := pages.NotFound()
		component.Render(r.Context(), w)
	})

	fmt.Printf("Server is running on port: %s", PORT)
	err := http.ListenAndServe(PORT, r)
	fmt.Printf("Server error %v", err)

	return err
}
