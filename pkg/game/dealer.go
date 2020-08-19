package game

import (
	"math/rand"
	"time"
)

// Dealer structure
type Dealer struct{}

// NewDealer creates a dealer
func NewDealer() Dealer {
	return Dealer{}
}

// Shuffle will shuffle deck
func (d Dealer) Shuffle(deck IDeck) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(deck.Count(), func(i, j int) {
		deck.Swap(i, j)
	})
}

// Deal will deal out deck
func (d *Dealer) Deal(deck IDeck, players IPlayerCollection) {
	for i := 0; i < deck.Count(); i++ {
		card, _ := deck.Next()
		players.Next().Hand.Add(card)
	}
}
