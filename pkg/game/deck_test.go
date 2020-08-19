package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDeck struct {
	Tracker map[string]int
	Length  int
}

func NewMockDeck(length int) MockDeck {
	return MockDeck{
		Tracker: make(map[string]int),
		Length:  length,
	}
}

func (m *MockDeck) Swap(i, j int) {
	m.Tracker["Swap"]++
}

func (m *MockDeck) Count() int {
	m.Tracker["Count"]++
	return m.Length
}

func (m *MockDeck) Next() (card Card, err error) {
	current := m.Tracker["Next"]
	m.Tracker["Next"]++
	if current < m.Length {
		return Card{}, fmt.Errorf("No More Cards")
	}
	return Card{current}, nil
}

func (m *MockDeck) Add(card Card) {
	m.Tracker["Add"]++
}

func TestNewFullDeck(t *testing.T) {
	deck := NewFullDeck()
	assert.Equal(t, 44, len(deck.Cards))
}

func TestSwap(t *testing.T) {
	deck := Deck{Cards: []Card{{1}, {2}, {3}, {4}}}

	assert.Equal(t, 2, deck.Cards[1].Value)
	assert.Equal(t, 3, deck.Cards[2].Value)

	deck.Swap(1, 2)

	assert.Equal(t, 3, deck.Cards[1].Value)
	assert.Equal(t, 2, deck.Cards[2].Value)
}

func TestValidNext(t *testing.T) {
	deck := Deck{Cards: []Card{{1}, {3}, {4}, {2}}}
	var tests = []struct {
		count int
	}{
		{1},
		{3},
		{4},
		{2},
	}
	for _, test := range tests {
		card, _ := deck.Next()
		assert.Equal(t, test.count, card.Value)
	}
}

func TestInvalidNext(t *testing.T) {
	deck := Deck{Cards: []Card{{1}}}
	card, noerr := deck.Next()
	assert.Equal(t, card.Value, 1)
	assert.Equal(t, nil, noerr)
	_, err := deck.Next()
	expectedErr := fmt.Errorf("No More Cards")
	assert.Equal(t, expectedErr, err)
}

func TestCount(t *testing.T) {
	var tests = []struct {
		count int
	}{
		{0},
		{1},
		{2},
	}
	for _, test := range tests {
		deck := Deck{}
		for i := 0; i < test.count; i++ {
			deck.Cards = append(deck.Cards, Card{})
		}
		assert.Equal(t, int(test.count), deck.Count())
	}
}

func TestAdd(t *testing.T) {
	deck := Deck{}
	var tests = []struct {
		count int
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
