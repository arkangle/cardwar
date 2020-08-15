package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, 44, len(deck.Cards))
}

func TestShuffle(t *testing.T) {
	deck := Deck{Cards: []Card{{1}, {2}, {3}, {4}}}
	deck.Shuffle()
	assert.Equal(t, 4, len(deck.Cards))
}

func TestValidDealCard(t *testing.T) {
	deck := Deck{Cards: []Card{{1}, {2}, {3}, {4}}}
	var tests = []struct {
		index int
		value int8
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
	}
	for _, test := range tests {
		deck.Index = test.index
		card, _ := deck.DealCard()
		assert.Equal(t, test.value, card.Value)
	}
}

func TestInvalidDealCard(t *testing.T) {
	deck := Deck{
		Cards: []Card{{1}},
	}
	card, noerr := deck.DealCard()
	assert.Equal(t, card.Value, int8(1))
	assert.Equal(t, nil, noerr)
	_, err := deck.DealCard()
	expectedErr := fmt.Errorf("No More Cards")
	assert.Equal(t, expectedErr, err)
}
