package game

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck structure
type Deck struct {
	Cards []Card
	Index int
}

// NewDeck creates Deck with cards
func NewDeck() Deck {
	cards := []Card{}
	for num := 1; num < 12; num++ {
		for count := 0; count < 4; count++ {
			card, _ := NewCard(num)
			cards = append(cards, card)
		}
	}
	return Deck{
		Cards: cards,
	}
}

// Shuffle will return shuffled deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

// DealCard will advance index and return it, maybe error
func (d *Deck) DealCard() (Card, error) {
	if d.Index == d.count() {
		return Card{}, fmt.Errorf("No More Cards")
	}
	card := d.Cards[d.Index]
	d.Index++

	return card, nil
}

func (d Deck) count() int {
	return len(d.Cards)
}
