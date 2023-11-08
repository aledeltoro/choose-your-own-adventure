package templating

import (
	"html/template"
	"io"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadTemplates(t *testing.T) {
	c := require.New(t)

	workingDir, err := os.Getwd()
	c.NoError(err)

	workingDir = strings.ReplaceAll(workingDir, "internal/templating", "")

	err = os.Chdir(workingDir)
	c.NoError(err)

	templates, err := LoadTemplates()
	c.NoError(err)
	c.NotEmpty(templates)
}

func TestRenderTemplate(t *testing.T) {
	c := require.New(t)

	recorder := httptest.NewRecorder()

	sampleTemplate := `
	{{define "base"}}
	<h1>Hello World</h1>
	{{end}}
	`

	tmpl := template.New("test")
	tmpl = template.Must(tmpl.Parse(sampleTemplate))

	RenderTemplate(recorder, tmpl, nil)

	data, err := parseResponse(recorder)
	c.NoError(err)

	c.Contains(string(data), `<h1>Hello World</h1>`)
}

func TestRenderTemplateError(t *testing.T) {
	c := require.New(t)

	recorder := httptest.NewRecorder()

	invalidTemplate := `<h1>Hello World</h1>`

	tmpl := template.New("test")
	tmpl = template.Must(tmpl.Parse(invalidTemplate))

	RenderTemplate(recorder, tmpl, nil)

	data, err := parseResponse(recorder)
	c.NoError(err)

	c.Contains(string(data), "Unexpected server error")
}

func parseResponse(recorder *httptest.ResponseRecorder) ([]byte, error) {
	result := recorder.Result()
	defer result.Body.Close()

	return io.ReadAll(result.Body)
}
