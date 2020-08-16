package game

// Dealer structure
type Dealer struct {
	Deck    Deck
	Players PlayerCollection
}

// NewDealer creates a dealer
func NewDealer(players PlayerCollection, deck Deck) Dealer {
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
	for card, err := d.Deck.Deal(); err == nil; card, err = d.Deck.Deal() {
		d.Players.Next().Hand.Add(card)
	}
}
