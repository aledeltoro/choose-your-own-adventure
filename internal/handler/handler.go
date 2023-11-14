package handler

import (
	"net/http"
	"strings"

	"github.com/aledeltoro/choose-your-own-adventure/internal/service"
	"github.com/aledeltoro/choose-your-own-adventure/internal/templating"
)

// Handler processes requests to render adventure story
type Handler struct {
	service service.StoryService
}

// NewHandler creates a handler instance
func NewHandler(service service.StoryService) Handler {
	return Handler{
		service: service,
	}
}

// HandleRenderStory renders the corresponding story page
func (h Handler) HandleRenderStory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentStory := strings.ReplaceAll(r.URL.Path, "/cyoa/", "")

		tmpl, input := h.service.GetStory(currentStory)

		templating.RenderTemplate(w, tmpl, input)
	}
}
