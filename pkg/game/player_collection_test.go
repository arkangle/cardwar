package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayerCollection(t *testing.T) {
	names := []string{"Player1", "Player2", "Player3"}
	playerCollection := NewPlayerCollection(names)
	assert.Equal(t, 3, playerCollection.Count())
	for i, name := range names {
		assert.Equal(t, name, playerCollection.Players[i].Name)
	}
}

func TestNextPlayer(t *testing.T) {
	names := []string{"Player1", "Player2", "Player3"}
	players := NewPlayerCollection(names)
	for i := 0; i < 7; i++ {
		index := i % len(names)
		assert.Equal(t, players.Next().Name, names[index])
	}
}
