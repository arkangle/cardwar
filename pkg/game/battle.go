package game

import "fmt"

// Battle structure
type Battle struct {
	Players []Player
}

// NewBattle creates new Battle
func NewBattle(players []Player) (Battle, error) {
	if len(players) < 2 {
		err := fmt.Errorf("Must have at least 2 players")
		return Battle{}, err
	}
	return Battle{Players: players}, nil
}
