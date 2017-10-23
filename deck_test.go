package fivetwo_test

import (
	"testing"

	. "github.com/praffn/fivetwo"
)

func TestNewStandardDeckGiven0JokersHasLeh52(t *testing.T) {
	deck := NewStandardDeck(0)
	if len(deck) != 52 {
		t.Errorf("Expected 52 got %d", len(deck))
	}
}

func TestNewStandardDeckGiven2JokersHasLen54(t *testing.T) {
	deck := NewStandardDeck(2)
	if len(deck) != 54 {
		t.Errorf("Expected 54 got %d", len(deck))
	}
}

func TestShuffle(t *testing.T) {
	// this test is kind of stupid... it can fail
	// even if shuffle went good but... oh well...
	// it is highly unlikely
	deck := NewStandardDeck(0)
	card1 := deck[10]
	deck.Shuffle()
	card2 := deck[10]
	if card1 == card2 {
		t.Fail()
	}
}

func TestLolMethod(t *testing.T) {
	deck := NewStandardDeck(0)
	bytes := deck.Bytes()
	dfb := DeckFromBytes(bytes)
	if len(deck) != len(dfb) {
		t.Errorf("Expected len of decks to be equal. %d != %d", len(deck), len(dfb))
		return
	}
	for i := range deck {
		if deck[i] != dfb[i] {
			t.Errorf("Expected card at position %d to be same. Expected %v, got %v", i, deck[i], dfb[i])
		}
	}
}
