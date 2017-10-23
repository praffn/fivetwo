package fivetwo

// Suit is the cards suit (sometimes refered to as colour)
type Suit uint8

// Suits
const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

// SuitNames maps a suit to its string name
var SuitNames = map[Suit]string{
	Clubs:    "Clubs",
	Diamonds: "Diamonds",
	Hearts:   "Hearts",
	Spades:   "Spades",
}

// A SuitComparator is a function that compares suits
type SuitComparator func(Suit, Suit) int

// DefaultSuitComparator compares by value
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
