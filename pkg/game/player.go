package game

// Player structure
type Player struct {
	Name string
	Hand Deck
}

// NewPlayer creates new Player
func NewPlayer(name string) Player {
	return Player{
		Name: name,
		Hand: Deck{},
	}
}
