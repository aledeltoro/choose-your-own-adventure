package templating

import (
	"fmt"
	"html/template"
	"path/filepath"
)

const (
	templateLayoutsPath = "templates/layout/*.html"
	templatesPath       = "templates/*.html"
)

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
