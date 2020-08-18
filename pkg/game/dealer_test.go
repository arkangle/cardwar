package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDealer(t *testing.T) {
	deck := Deck{Cards: []Card{{1}, {2}, {3}}}
	names := []string{"Player1", "Player2", "Player3"}
	players := NewPlayerCollection(names)
	dealer := NewDealer(&players, &deck)
	assert.Equal(t, &deck, dealer.Deck)
	assert.Equal(t, &players, dealer.Players, nil)
}

func TestDealerShuffle(t *testing.T) {
	mockDeck := NewMockDeck(5)
	mockPlayerCollection := NewMockPlayerCollection("Player", 2)
	dealer := Dealer{
		Players: &mockPlayerCollection,
		Deck:    &mockDeck,
	}

	assert.Equal(t, 0, mockDeck.Tracker["Shuffle"])
	dealer.Shuffle()
	assert.Equal(t, 1, mockDeck.Tracker["Shuffle"])
}

func TestDealerDeal(t *testing.T) {
	count := 5
	mockDeck := NewMockDeck(count)
	mockPlayerCollection := NewMockPlayerCollection("Player", count)
	dealer := Dealer{
		Players: &mockPlayerCollection,
		Deck:    &mockDeck,
	}
	dealer.Deal()
	assert.Equal(t, count, mockDeck.Tracker["Next"])
	assert.Equal(t, count+1, mockDeck.Tracker["Count"])
	assert.Equal(t, count, mockPlayerCollection.Tracker["Next"])
}
