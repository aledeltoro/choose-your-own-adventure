package handler

import (
	"gophercises/choose-your-own-adventure/internal/models"
	"gophercises/choose-your-own-adventure/internal/templating"
	"html/template"
	"net/http"
	"strings"
)

// Handler represents the interface to handle the adventure story rendering
type Handler interface {
	HandleRenderStory() http.HandlerFunc
}

type handler struct {
	templates map[string]*template.Template
	input     map[string]models.AdventureInput
}

// NewHandler creates a handler instance
func NewHandler(templates map[string]*template.Template, input map[string]models.AdventureInput) Handler {
	return handler{
		templates: templates,
		input:     input,
	}
}

// HandleRenderStory renders the corresponding story page
func (h handler) HandleRenderStory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentStory := strings.ReplaceAll(r.URL.Path, "/cyoa/", "")

		if currentStory == "" {
			templating.RenderTemplate(w, h.templates["story.html"], h.input["intro"])
			return
		}

		nextStory, ok := h.input[currentStory]
		if !ok {
			templating.RenderTemplate(w, h.templates["not_found.html"], nil)
			return
		}

		templating.RenderTemplate(w, h.templates["story.html"], nextStory)
	}
}
