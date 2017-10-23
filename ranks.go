package fivetwo

// Rank is the value of a card
type Rank uint8

// All defined ranks. Joker after ace after king
const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	Joker
)

// RankNames maps a rank to its string name
var RankNames = map[Rank]string{
	Two:   "Two",
	Three: "Three",
	Four:  "Four",
	Five:  "Five",
	Six:   "Six",
	Seven: "Seven",
	Eight: "Eight",
	Nine:  "Nine",
	Ten:   "Ten",
	Jack:  "Jack",
	Queen: "Queen",
	King:  "King",
	Ace:   "Ace",
	Joker: "Joker",
}

// A RankComparator is a function that compares ranks
type RankComparator func(Rank, Rank) int

// DefaultRankComparator compares by value
func DefaultRankComparator(a, b Rank) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
