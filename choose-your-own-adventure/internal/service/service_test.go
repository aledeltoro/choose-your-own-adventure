package service

import (
	"gophercises/choose-your-own-adventure/internal/models"
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStoryIntro(t *testing.T) {
	c := require.New(t)

	templates := map[string]*template.Template{
		"story.html": template.New("story"),
	}

	adventure := map[string]*models.AdventureInput{
		"intro": {
			Title: "The Little Blue Gopher",
			Story: []string{"Testing"},
		},
	}

	service := NewStoryService(templates, adventure)

	tmpl, input := service.GetStory("")
	c.Equal(template.New("story"), tmpl)
	c.Equal(adventure["intro"], input)
}

func TestGetStoryNextStory(t *testing.T) {
	c := require.New(t)

	templates := map[string]*template.Template{
		"story.html": template.New("story"),
	}

	adventure := map[string]*models.AdventureInput{
		"intro": {
			Title: "The Little Blue Gopher",
			Story: []string{"Testing"},
		},
		"new-york": {
			Title: "Visiting New York",
			Story: []string{"Testing"},
		},
	}

	service := NewStoryService(templates, adventure)

	tmpl, input := service.GetStory("new-york")
	c.Equal(template.New("story"), tmpl)
	c.Equal(adventure["new-york"], input)
}

func TestGetStoryNotFound(t *testing.T) {
	c := require.New(t)

	templates := map[string]*template.Template{
		"story.html":     template.New("story"),
		"not_found.html": template.New("not_found"),
	}

	adventure := map[string]*models.AdventureInput{}

	service := NewStoryService(templates, adventure)

	tmpl, input := service.GetStory("invalid")
	c.Equal(template.New("not_found"), tmpl)
	c.Nil(input)
}
