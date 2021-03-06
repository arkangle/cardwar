package game

import "fmt"

// Battle structure
type Battle struct {
	Players IPlayerCollection
}

// NewBattle creates new Battle
func NewBattle(players IPlayerCollection) (Battle, error) {
	if players.Count() < 2 {
		err := fmt.Errorf("Must have at least 2 players")
		return Battle{}, err
	}
	return Battle{Players: players}, nil
}
