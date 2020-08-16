package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	name := "Name"
	player := NewPlayer(name)
	assert.Equal(t, name, player.Name)
}
