package models

import "encoding/json"

type adventureInput struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func NewAdventureInput(data []byte) (map[string]adventureInput, error) {
	var input map[string]adventureInput

	err := json.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}

	return input, nil
}
