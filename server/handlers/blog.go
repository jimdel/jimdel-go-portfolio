package handlers

import (
	"fmt"
	"jimdel/pkg/server/helpers"
	"jimdel/pkg/web/views/pages"
	"net/http"
	"sort"

	"github.com/go-chi/chi"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	articles := getAllBlogPosts()
	articleId := chi.URLParam(r, "articleId")
	if articleId == "" {
		pages.Blog(articles).Render(r.Context(), w)
	}

	for _, article := range articles {
		if article.Meta.Title == articleId {
			pages.Article(article).Render(r.Context(), w)
			return
		}
	}

}
func getAllBlogPosts() []helpers.Markdown {
	files := helpers.GetAllFilesInDir("blog")
	articles := []helpers.Markdown{}
	for _, file := range files {
		fmt.Println(file)
		fileName := fmt.Sprintf("blog/%s", file)
		md := NewMarkdown(fileName)
		articles = append(articles, md)
	}
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Meta.Date > articles[j].Meta.Date
	})
	return articles
}
