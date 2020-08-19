package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNewBattle(t *testing.T) {
	playerNames := []string{"Player1", "Player2"}
	players := NewPlayerCollection(playerNames)
	battle, err := NewBattle(&players)
	assert.Equal(t, battle.Players.Next().Name, playerNames[0])
	assert.Equal(t, battle.Players.Next().Name, playerNames[1])
	assert.Equal(t, err, nil)
}

func TestInvalidNewBattle(t *testing.T) {
	playerNames := []string{"Player1"}
	players := NewPlayerCollection(playerNames)
	_, err := NewBattle(&players)
	assert.EqualError(t, err, "Must have at least 2 players")
}
