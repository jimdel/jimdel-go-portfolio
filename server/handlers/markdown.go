package handlers

import (
	"bytes"
	"jimdel/pkg/server/helpers"
	"log"
	"os"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/google/uuid"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// MarkdownMetaData represents the frontmatter structure

func NewMarkdown(path string) helpers.Markdown {
	// Read file
	cwd, _ := os.Getwd()
	file, err := os.ReadFile(cwd + path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse frontmatter
	var metadata helpers.MarkdownMetaData
	markdown, err := frontmatter.Parse(bytes.NewReader(file), &metadata)
	if err != nil {
		// Handle files without frontmatter
		markdown = file
		metadata = helpers.MarkdownMetaData{
			Title: "Untitled",
			Date:  "No date",
			Blurb: extractBlurb(string(file), 250),
		}
	}

	// Configure goldmark with extensions
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,            // GitHub Flavored Markdown
			extension.Table,          // Tables
			extension.Strikethrough,  // Strikethrough
			extension.TaskList,       // Task lists
			extension.DefinitionList, // Definition lists
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // Auto generate heading IDs
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(), // Allow raw HTML
		),
	)

	// Convert markdown to HTML
	var buf bytes.Buffer
	if err := md.Convert(markdown, &buf); err != nil {
		log.Fatalf("Error converting markdown: %v", err)
	}

	// If no blurb in frontmatter, generate one
	if metadata.Blurb == "" {
		metadata.Blurb = extractBlurb(string(markdown), 250)
	}

	return helpers.Markdown{
		Html: buf.String(),
		Meta: metadata,
		UUID: uuid.New().String(),
	}
}

// extractBlurb generates a blurb from the content
func extractBlurb(content string, length int) string {
	// Remove any HTML tags
	content = strings.ReplaceAll(content, "\n", " ")

	// Trim to length
	if len(content) > length {
		content = content[:length] + "..."
	}

	return strings.TrimSpace(content)
}
