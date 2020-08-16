package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDealer(t *testing.T) {
	deck := Deck{Cards: []Card{{1}, {2}, {3}}}
	names := []string{"Player1", "Player2", "Player3"}
	players := NewPlayerCollection(names)
	dealer := NewDealer(players, deck)
	assert.Equal(t, deck, dealer.Deck)
	assert.Equal(t, players, dealer.Players)
}
