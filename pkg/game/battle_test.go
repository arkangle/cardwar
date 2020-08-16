package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNewBattle(t *testing.T) {
	players := []Player{
		Player{Name: "Player1"},
		Player{Name: "Player2"},
	}
	battle, err := NewBattle(players)
	assert.Equal(t, battle.Players[0].Name, "Player1")
	assert.Equal(t, battle.Players[1].Name, "Player2")
	assert.Equal(t, err, nil)
}

func TestInvalidNewBattle(t *testing.T) {
	players := []Player{
		Player{Name: "Player1"},
	}
	_, err := NewBattle(players)
	assert.EqualError(t, err, "Must have at least 2 players")
}
