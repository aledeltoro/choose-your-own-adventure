package handler

import (
	"fmt"
	"gophercises/choose-your-own-adventure/internal/models"
	"html/template"
	"log"
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

		var err error

		if currentStory == "" {
			err = h.templates["story.html"].ExecuteTemplate(w, "base", h.input["intro"])
			if err != nil {
				// Respond with error message to page
				log.Println("executing template failed: ", err)
				fmt.Fprintln(w, "Unexpected error")
				return
			}

			return
		}

		// STRUCTURE THE HTML BETTER THEN ADD CSS

		nextStory, ok := h.input[currentStory]
		if !ok {
			// Response with 404
			err := h.templates["not_found.html"].ExecuteTemplate(w, "base", nil)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = h.templates["story.html"].ExecuteTemplate(w, "base", nextStory)
		if err != nil {
			log.Println("executing template failed: ", err)
			fmt.Fprintln(w, "Unexpected error")
		}
	}
}
