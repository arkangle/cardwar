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

func TestValidDeal(t *testing.T) {
	deck := Deck{Cards: []Card{{1}, {3}, {4}, {2}}}
	var tests = []struct {
		count int8
	}{
		{1},
		{3},
		{4},
		{2},
	}
	for _, test := range tests {
		card, _ := deck.Deal()
		assert.Equal(t, test.count, card.Value)
	}
}

func TestInvalidDeal(t *testing.T) {
	deck := Deck{Cards: []Card{{1}}}
	card, noerr := deck.Deal()
	assert.Equal(t, card.Value, int8(1))
	assert.Equal(t, nil, noerr)
	_, err := deck.Deal()
	expectedErr := fmt.Errorf("No More Cards")
	assert.Equal(t, expectedErr, err)
}

func TestCount(t *testing.T) {
	var tests = []struct {
		count int8
	}{
		{0},
		{1},
		{2},
	}
	for _, test := range tests {
		deck := Deck{}
		for i := int8(0); i < test.count; i++ {
			deck.Cards = append(deck.Cards, Card{})
		}
		assert.Equal(t, int(test.count), deck.Count())
	}
}

func TestAdd(t *testing.T) {
	deck := Deck{}
	var tests = []struct {
		count int8
	}{
		{1},
		{2},
		{3},
		{4},
	}
	for _, test := range tests {
		deck.Add(Card{test.count})
		assert.Equal(t, int(test.count), len(deck.Cards))
	}
}
