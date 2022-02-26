package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCharacter(t *testing.T) {
	// Initialization
	name := "Santiago Leira"
	films := []string{"film1", "film2", "film3"}

	// Operation
	response := NewCharacter(name, films)

	// Validation
	assert.Equal(t, name, response.Name)
	assert.Equal(t, films, response.Films)
}
