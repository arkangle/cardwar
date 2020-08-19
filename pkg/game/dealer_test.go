package game

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDealer(t *testing.T) {
	dealer := NewDealer()
	assert.Equal(t, reflect.TypeOf(Dealer{}), reflect.TypeOf(dealer))
}

func TestDealerShuffle(t *testing.T) {
	mockDeck := NewMockDeck(5)
	dealer := Dealer{}

	dealer.Shuffle(&mockDeck)
	assert.Equal(t, 4, mockDeck.Tracker["Swap"])
}

func TestDealerDeal(t *testing.T) {
	count := 5
	mockDeck := NewMockDeck(count)
	mockPlayerCollection := NewMockPlayerCollection("Player", count)
	dealer := Dealer{}
	dealer.Deal(&mockDeck, &mockPlayerCollection)
	assert.Equal(t, count, mockDeck.Tracker["Next"])
	assert.Equal(t, count+1, mockDeck.Tracker["Count"])
	assert.Equal(t, count, mockPlayerCollection.Tracker["Next"])
}
