package game

// Dealer structure
type Dealer struct {
	Deck    IDeck
	Players IPlayerCollection
}

// NewDealer creates a dealer
func NewDealer(players IPlayerCollection, deck IDeck) Dealer {
	return Dealer{
		Deck:    deck,
		Players: players,
	}
}

// Shuffle will shuffle deck
func (d *Dealer) Shuffle() {
	d.Deck.Shuffle()
}

// Deal will deal out deck
func (d *Dealer) Deal() {
	for i := 0; i < d.Deck.Count(); i++ {
		card, _ := d.Deck.Next()
		d.Players.Next().Hand.Add(card)
	}
}
