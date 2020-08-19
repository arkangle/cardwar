package game

import (
	"fmt"
)

// IDeck interface for Deck
type IDeck interface {
	Swap(int, int)
	Next() (Card, error)
	Count() int
	Add(Card)
}

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

// Next will advance index and return it, maybe error
func (d *Deck) Next() (Card, error) {
	if d.Count() == 0 {
		return Card{}, fmt.Errorf("No More Cards")
	}
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card, nil
}

// Count cards
func (d Deck) Count() int {
	return len(d.Cards)
}

// Add card to deck
func (d *Deck) Add(card Card) {
	d.Cards = append(d.Cards, card)
}

// Swap cards
func (d *Deck) Swap(i, j int) {
	d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
}
