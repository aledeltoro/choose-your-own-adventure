package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAdventureInput(t *testing.T) {
	c := require.New(t)

	data, err := os.ReadFile("../../gopher.json")
	c.NoError(err)

	input, err := NewAdventureInput(data)
	c.NoError(err)
	c.NotEmpty(input)
}
