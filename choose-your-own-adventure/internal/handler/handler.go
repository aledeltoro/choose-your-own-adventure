package handler

import (
	"net/http"
	"strings"

	"github.com/aledeltoro/choose-your-own-adventure/internal/service"
	"github.com/aledeltoro/choose-your-own-adventure/internal/templating"
)

// Handler represents the interface to handle the adventure story rendering
type Handler interface {
	HandleRenderStory() http.HandlerFunc
}

type handler struct {
	service service.StoryService
}

// NewHandler creates a handler instance
func NewHandler(service service.StoryService) Handler {
	return handler{
		service: service,
	}
}

// HandleRenderStory renders the corresponding story page
func (h handler) HandleRenderStory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentStory := strings.ReplaceAll(r.URL.Path, "/cyoa/", "")

		tmpl, input := h.service.GetStory(currentStory)

		templating.RenderTemplate(w, tmpl, input)
	}
}
