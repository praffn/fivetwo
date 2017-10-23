package fivetwo

// A Card is a rank and a suit
type Card struct {
	Rank Rank
	Suit Suit
}

// String returns a string representation of the card
// in the form of "<Rank> of <Suit>", e.g. King of Spades
func (c Card) String() string {
	if c.Rank == Joker {
		return "Joker"
	}
	return RankNames[c.Rank] + " of " + SuitNames[c.Suit]
}

// Byte returns a byte representation of the card.
// The two least significant bits will be the suit
// and the next four bits will be the rank.
func (c Card) Byte() (b byte) {
	b = byte(c.Suit)
	b |= byte(c.Rank) << 2
	return
}

// CardFromByte returns a card from the given byte.
func CardFromByte(b byte) (card Card) {
	card.Suit = Suit(b & 3)
	card.Rank = Rank(b >> 2)
	return card
}
