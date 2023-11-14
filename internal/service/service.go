package service

import (
	"html/template"

	"github.com/aledeltoro/choose-your-own-adventure/internal/models"
)

// StoryService service to handle story rendering
type StoryService struct {
	templates map[string]*template.Template
	adventure map[string]*models.AdventureInput
}

// NewStoryService constructor to create storyService instance
func NewStoryService(templates map[string]*template.Template, adventure map[string]*models.AdventureInput) StoryService {
	return StoryService{
		templates: templates,
		adventure: adventure,
	}
}

// GetStory service that returns the next story
func (s StoryService) GetStory(currentStory string) (*template.Template, *models.AdventureInput) {
	if currentStory == "" {
		return s.templates["story.html"], s.adventure["intro"]
	}

	nextStory, ok := s.adventure[currentStory]
	if !ok {
		return s.templates["not_found.html"], nil
	}

	return s.templates["story.html"], nextStory
}
