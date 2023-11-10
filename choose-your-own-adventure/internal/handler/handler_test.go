package handler

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aledeltoro/gophercises/choose-your-own-adventure/internal/models"
	"github.com/aledeltoro/gophercises/choose-your-own-adventure/internal/service"
	"github.com/stretchr/testify/require"
)

func TestHandlerRenderStory(t *testing.T) {
	c := require.New(t)

	sampleTemplate := `
	{{define "base"}}
	<h1>{{.Title}}</h1>
	{{end}}
	`

	tmpl := template.New("test")
	tmpl = template.Must(tmpl.Parse(sampleTemplate))

	templates := map[string]*template.Template{
		"story.html": tmpl,
	}

	adventure := map[string]*models.AdventureInput{
		"intro": {
			Title: "The Little Blue Gopher",
		},
	}

	service := service.NewStoryService(templates, adventure)

	handler := NewHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/cyoa/", nil)
	w := httptest.NewRecorder()

	handler.HandleRenderStory()(w, req)

	resp := w.Result()

	defer func() {
		_ = resp.Body.Close()
	}()

	data, err := io.ReadAll(resp.Body)
	c.NoError(err)

	c.Equal(http.StatusOK, resp.StatusCode)
	c.Contains(string(data), `<h1>The Little Blue Gopher</h1>`)
}

func TestHandlerRenderStoryChapterNotFound(t *testing.T) {
	c := require.New(t)

	sampleTemplate := `
	{{define "base"}}
	<h1>Chapter not found</h1>
	{{end}}
	`

	tmpl := template.New("test")
	tmpl = template.Must(tmpl.Parse(sampleTemplate))

	templates := map[string]*template.Template{
		"not_found.html": tmpl,
	}

	adventure := map[string]*models.AdventureInput{}

	service := service.NewStoryService(templates, adventure)

	handler := NewHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/cyoa/not-found", nil)
	w := httptest.NewRecorder()

	handler.HandleRenderStory()(w, req)

	resp := w.Result()

	defer func() {
		_ = resp.Body.Close()
	}()

	data, err := io.ReadAll(resp.Body)
	c.NoError(err)

	c.Contains(string(data), `<h1>Chapter not found</h1>`)
}
