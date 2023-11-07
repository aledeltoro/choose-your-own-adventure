package handler

import (
	"gophercises/choose-your-own-adventure/internal/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Handler interface {
	HandleRenderStory() http.HandlerFunc
}

type handler struct {
	template *template.Template
	input    map[string]models.AdventureInput
}

// NewHandler creates a handler instance
func NewHandler(template *template.Template, input map[string]models.AdventureInput) Handler {
	return handler{
		template: template,
		input:    input,
	}
}

// HandleRenderStory renders the corresponding story page
func (h handler) HandleRenderStory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentStory := strings.ReplaceAll(r.URL.Path, "/cyoa/", "")

		var err error

		if currentStory == "" {
			err = h.template.Execute(w, h.input["intro"])
			if err != nil {
				// Respond with error message to page
				log.Println("executing template failed: ", err)
				return
			}
		}

		err = h.template.Execute(w, h.input[currentStory])
		if err != nil {
			log.Println("executing template failed: ", err)
		}
	}
}
