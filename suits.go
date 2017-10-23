package fivetwo

type Suit uint8

// Suits
const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

var SuitNames = map[Suit]string{
	Clubs:    "Clubs",
	Diamonds: "Diamonds",
	Hearts:   "Hearts",
	Spades:   "Spades",
}

type SuitComparator func(Suit, Suit) int

func DefaultSuitComparator(a, b Suit) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
