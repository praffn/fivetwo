package fivetwo_test

import (
	"testing"

	. "github.com/praffn/fivetwo"
)

func TestString(t *testing.T) {
	card := Card{Rank: King, Suit: Spades}
	if card.String() != "King of Spades" {
		t.Fail()
	}
}

func TestStringJokerPrintsOnlyJoker(t *testing.T) {
	card := Card{Rank: Joker, Suit: Hearts}
	if card.String() != "Joker" {
		t.Errorf("Expected 'Joker' got %s", card.String())
	}
}

func TestByte(t *testing.T) {
	card := Card{Rank: King, Suit: Spades}
	if card.Byte() != 0x2f {
		t.Errorf("Expected 37, got %x", card.Byte())
	}
}

func TestCardFromByte(t *testing.T) {
	card := CardFromByte(0x2f)
	if card.String() != "King of Spades" {
		t.Errorf("Expected King of Spades, got %s", card.String())
	}
}
