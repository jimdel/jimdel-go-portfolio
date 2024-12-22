package handlers

import (
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type MetaData struct {
	Title string
	Date  string
}

type Markdown struct {
	Meta MetaData
	Html string
}

// TODO: add support for Front Matter
func NewMarkdown(path string) Markdown {

	cwd, _ := os.Getwd()
	file, err := os.ReadFile(cwd + path)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.FencedCode | parser.Tables | parser.Strikethrough | parser.SpaceHeadings | parser.HeadingIDs
	p := parser.NewWithExtensions(extensions)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	html := markdown.ToHTML(file, p, renderer)

	return Markdown{
		Html: string(html),
		Meta: MetaData{
			Title: "Test",
			Date:  "2042-12-22",
		},
	}
}
