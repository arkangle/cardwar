package game

import "fmt"

// Card structure
type Card struct {
	Value int8
}

// NewCard creates Card with value
func NewCard(value int) (Card, error) {
	if value < 1 || value > 11 {
		return Card{}, fmt.Errorf("Card value (%d) must between 1 and 11", value)
	}
	return Card{Value: int8(value)}, nil
}
