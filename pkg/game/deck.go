package game

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck structure
type Deck struct {
	Cards []Card
}

// NewFullDeck creates Deck with cards
func NewFullDeck() Deck {
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

// Next will advance index and return it, maybe error
func (d *Deck) Next() (Card, error) {
	if d.Count() == 0 {
		return Card{}, fmt.Errorf("No More Cards")
	}
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card, nil
}

// Count will count cards
func (d Deck) Count() int {
	return len(d.Cards)
}

// Add adds card to deck
func (d *Deck) Add(card Card) {
	d.Cards = append(d.Cards, card)
}
