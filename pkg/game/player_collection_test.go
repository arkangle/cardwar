package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockPlayerCollection struct {
	Tracker map[string]int
	Player  Player
	Length  int
}

func NewMockPlayerCollection(name string, length int) MockPlayerCollection {
	return MockPlayerCollection{
		Tracker: make(map[string]int),
		Player:  NewPlayer(name + string(length)),
		Length:  length,
	}
}

func (m *MockPlayerCollection) Count() int {
	m.Tracker["Count"]++
	return m.Length
}

func (m *MockPlayerCollection) Next() *Player {
	m.Tracker["Next"]++
	return &m.Player
}

func TestNewPlayerCollection(t *testing.T) {
	names := []string{"Player1", "Player2", "Player3"}
	playerCollection := NewPlayerCollection(names)
	assert.Equal(t, 3, playerCollection.Count())
	for i, name := range names {
		assert.Equal(t, name, playerCollection.Players[i].Name)
	}
}

func TestNextPlayer(t *testing.T) {
	names := []string{"Player1", "Player2", "Player3"}
	players := NewPlayerCollection(names)
	for i := 0; i < 7; i++ {
		index := i % len(names)
		assert.Equal(t, players.Next().Name, names[index])
	}
}
