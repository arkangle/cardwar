package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNewCard(t *testing.T) {
	var tests = []struct {
		given int
		value int
	}{
		{1, 1},
		{6, 6},
		{11, 11},
	}
	for _, test := range tests {
		card, _ := NewCard(test.given)
		assert.Equal(t, test.value, card.Value)
	}
}

func TestInvalidNewCard(t *testing.T) {
	var tests = []struct {
		given int
		value int
	}{
		{0, 0},
		{12, 12},
	}
	for _, test := range tests {
		expectedErr := fmt.Errorf("Card value (%d) must between 1 and 11", test.value)
		_, err := NewCard(test.given)
		assert.Equal(t, err, expectedErr)
	}
}
