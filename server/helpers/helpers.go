package helpers

import (
	"io/fs"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	"jimdel/pkg/content"

	"github.com/go-chi/chi"
	"gopkg.in/yaml.v3"
)

const (
	EmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func IsValidRegex(str string, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}

type Project struct {
	Title       string `yaml:"title"`
	URL         string `yaml:"url"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`
	CoverImage  string `yaml:"coverImg"`
}

type BlogPost struct {
	Title   string `yaml:"title"`
	Content string
}

type SiteConfig struct {
	ResumeURL string `yaml:"resume"`
	Socials   struct {
		LinkedIn string `yaml:"linkedin"`
		Github   string `yaml:"github"`
	}
	Bio string `yaml:"bio"`
}

type MarkdownMetaData struct {
	Title string `yaml:"title"`
	Date  string `yaml:"date"`
	Blurb string `yaml:"blurb"`
}

type Markdown struct {
	UUID string
	Meta MarkdownMetaData
	Html string
}

func ParseYaml[T any](fileName string, targetStruct *T) *T {
	yamlData, err := content.ContentFS.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	err = yaml.Unmarshal(yamlData, targetStruct)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return targetStruct
}

func GetAllFilesInDir(subDir string) []string {
	files, err := content.ContentFS.ReadDir(subDir)

	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	fileNames := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}

func GetSiteConfigValue(key string) string {
	siteConfig := &SiteConfig{}
	siteConfig = ParseYaml("site.config.yml", siteConfig)
	switch key {
	case "resume":
		return siteConfig.ResumeURL
	case "linkedin":
		return siteConfig.Socials.LinkedIn
	case "github":
		return siteConfig.Socials.Github
	case "bio":
		return siteConfig.Bio
	default:
		return ""
	}
}

func FileServer(r chi.Router, path string, root fs.FS) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(http.FS(root)))

		// Get the file path from the request
		filePath := strings.TrimPrefix(r.URL.Path, pathPrefix)

		// Detect content type based on file extension
		ext := filepath.Ext(filePath)
		contentType := mime.TypeByExtension(ext)
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}

		fs.ServeHTTP(w, r)
	})
}
