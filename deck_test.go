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

func TestBytes(t *testing.T) {
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

func TestDrawReturnsCard(t *testing.T) {
	needle := Card{Suit: Diamonds, Rank: Queen}
	cards := []Card{
		needle,
		Card{Suit: Spades, Rank: Eight},
	}
	deck := Deck(cards)
	card, err := deck.Draw()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	if card != needle {
		t.Errorf("Expected %v got %v", needle, card)
		return
	}
	if len(deck) != 1 {
		t.Errorf("Expected length of 1, got %d", len(deck))
	}
}

func TestDrawErrors(t *testing.T) {
	var deck Deck
	_, err := deck.Draw()
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestDrawBackReturnsCard(t *testing.T) {
	needle := Card{Suit: Diamonds, Rank: Queen}
	cards := []Card{
		Card{Suit: Spades, Rank: Eight},
		needle,
	}
	deck := Deck(cards)
	card, err := deck.DrawBack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	if card != needle {
		t.Errorf("Expected %v got %v", needle, card)
		return
	}
	if len(deck) != 1 {
		t.Errorf("Expected length of 1, got %d", len(deck))
	}
}

func TestDrawBackErrors(t *testing.T) {
	var deck Deck
	_, err := deck.DrawBack()
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestAdd(t *testing.T) {
	var deck Deck
	if len(deck) != 0 {
		t.Errorf("Expected empty deck")
		return
	}
	deck.Add(Card{})
	if len(deck) != 1 {
		t.Errorf("Expected deck with size 1, got %d", len(deck))
	}
}

func TestAddAddsToBack(t *testing.T) {
	cards := []Card{
		Card{Suit: Spades, Rank: Eight},
		Card{Suit: Diamonds, Rank: Ace},
	}
	deck := Deck(cards)
	deck.Add(Card{})
	card, err := deck.DrawBack()
	expected := Card{}
	if err != nil {
		t.Error(err)
		return
	}
	if card != expected {
		t.Errorf("Expected %v got %v", Card{}, expected)
	}
}

func TestAddFrontAddsToFront(t *testing.T) {
	cards := []Card{
		Card{Suit: Spades, Rank: Eight},
		Card{Suit: Diamonds, Rank: Ace},
	}
	deck := Deck(cards)
	deck.AddFront(Card{})
	card, err := deck.Draw()
	expected := Card{}
	if err != nil {
		t.Error(err)
		return
	}
	if card != expected {
		t.Errorf("Expected %v got %v", Card{}, expected)
	}
}

func TestPeekReturnsFrontCard(t *testing.T) {
	needle := Card{Suit: Diamonds, Rank: Queen}
	cards := []Card{
		needle,
		Card{Suit: Spades, Rank: Eight},
	}
	deck := Deck(cards)
	card, err := deck.Peek()
	if err != nil {
		t.Error(err)
		return
	}
	if card != needle {
		t.Errorf("Expected %v got %v", needle, card)
		return
	}
	if len(deck) != 2 {
		t.Errorf("Expected length of 2, got %d", len(deck))
		return
	}
}

func TestPeekBackReturnsBackCard(t *testing.T) {
	needle := Card{Suit: Diamonds, Rank: Queen}
	cards := []Card{
		Card{Suit: Spades, Rank: Eight},
		needle,
	}
	deck := Deck(cards)
	card, err := deck.PeekBack()
	if err != nil {
		t.Error(err)
		return
	}
	if card != needle {
		t.Errorf("Expected %v got %v", needle, card)
		return
	}
	if len(deck) != 2 {
		t.Errorf("Expected length of 2, got %d", len(deck))
		return
	}
}
