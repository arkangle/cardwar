package game

// PlayerCollection structure
type PlayerCollection struct {
	Players      []Player
	CurrentIndex int
}

// NewPlayerCollection creates new Player Collection
func NewPlayerCollection(names []string) PlayerCollection {
	players := []Player{}
	for _, name := range names {
		player := NewPlayer(name)
		players = append(players, player)
	}
	return PlayerCollection{
		Players: players,
	}
}

// Count will count players
func (c PlayerCollection) Count() int {
	return len(c.Players)
}

// Next will increment current and return Player
func (c *PlayerCollection) Next() *Player {
	index := c.CurrentIndex % c.Count()
	c.CurrentIndex++
	return &c.Players[index]
}
