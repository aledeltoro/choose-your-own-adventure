package templating

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const (
	templateLayoutsPath = "templates/layout/*.html"
	templatesPath       = "templates/*.html"
)

// RenderTemplate renders a HTML page
func RenderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Unexpected server error", 500)
		log.Println("executing template failed: %w", err)
	}
}

// LoadTemplates loads layout and page files
func LoadTemplates() (map[string]*template.Template, error) {
	templates := make(map[string]*template.Template)

	layoutFiles, err := filepath.Glob(templateLayoutsPath)
	if err != nil {
		return nil, fmt.Errorf("loading template layout files: %w", err)
	}

	pageFiles, err := filepath.Glob(templatesPath)
	if err != nil {
		return nil, fmt.Errorf("loading template files: %w", err)
	}

	for _, file := range pageFiles {
		filename := filepath.Base(file)
		files := append(layoutFiles, file)

		templates[filename] = template.Must(template.ParseFiles(files...))
	}

	return templates, nil
}
