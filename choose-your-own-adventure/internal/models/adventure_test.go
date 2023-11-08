package models

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAdventureInput(t *testing.T) {
	c := require.New(t)

	workingDir, err := os.Getwd()
	c.NoError(err)

	workingDir = strings.ReplaceAll(workingDir, "internal/models", "")

	err = os.Chdir(workingDir)
	c.NoError(err)

	data, err := os.ReadFile("assets/gopher.json")
	c.NoError(err)

	input, err := NewAdventureInput(data)
	c.NoError(err)
	c.NotEmpty(input)
}
