package models

import "encoding/json"

// AdventureInput input of the story and its options
type AdventureInput struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// NewAdventureInput parses input of the story and its options
func NewAdventureInput(data []byte) (map[string]AdventureInput, error) {
	var input map[string]AdventureInput

	err := json.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}

	return input, nil
}
