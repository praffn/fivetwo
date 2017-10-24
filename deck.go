package fivetwo

import (
	"bytes"
	"errors"
	"io"
	"math/rand"
)

// A Deck is a slice of cards
type Deck []Card

// NewStandardDeck returns a Deck with 52 different cards
// and the amount of jokers specified.
// No kind of order is guaranteed, but it is adviced to
// call shuffle on the deck before use.
func NewStandardDeck(jokers int) (d Deck) {
	for suit := range SuitNames {
		for rank := range RankNames {
			if rank == Joker {
				continue
			}
			d = append(d, Card{rank, suit})
		}
	}
	for jokers != 0 {
		jokers--
		d = append(d, Card{Joker, 0})
	}
	return
}

// Bytes returns a slice of bytes representing the deck
func (d Deck) Bytes() []byte {
	var bs bytes.Buffer
	for _, card := range d {
		err := bs.WriteByte(card.Byte())
		if err != nil {
			panic(err)
		}
	}
	return bs.Bytes()
}

// DeckFromBytes returns a Deck from the given byte slice
func DeckFromBytes(bs []byte) (d Deck) {
	reader := bytes.NewReader(bs)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		d = append(d, CardFromByte(b))
	}
}

// Shuffle used Fisher-Yates to shuffle the entire deck
func (d Deck) Shuffle() {
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

// Draw removes the card from the front of the deck and returns it
// An error will be returned if the deck is empty
func (d Deck) Draw() (c Card, err error) {
	if len(d) == 0 {
		err = errors.New("Deck is empty")
	} else {
		c, d = d[0], d[1:]
	}
	return
}

// DrawBack removes the card from the back of the deck and returns it.
// An error will be returned if the deck is empty
func (d Deck) DrawBack() (c Card, err error) {
	l := len(d)
	if l == 0 {
		err = errors.New("Deck is empty")
	} else {
		c, d = d[l-1], d[:l-1]
	}
	return
}

// Peek returns the card at the front of the deck without removing it
// An error will be returned if the deck is empty
func (d Deck) Peek() (c Card, err error) {
	if len(d) == 0 {
		err = errors.New("Deck is empty")
	} else {
		c = d[0]
	}
	return
}

// PeekBack returns the card at the back of the deck
// An error will be returned if the deck is empty
func (d Deck) PeekBack() (c Card, err error) {
	l := len(d)
	if l == 0 {
		err = errors.New("Deck is empty")
	} else {
		c = d[l-1]
	}
	return
}

// Add adds a card to the back of the deck
func (d *Deck) Add(c Card) {
	*d = append(*d, c)
}

// AddFront adds a card to the front of the deck
func (d *Deck) AddFront(c Card) {
	*d = append([]Card{c}, *d...)
}

// Empty returns true if deck is empty, otherwise false
func (d Deck) Empty() bool {
	return len(d) == 0
}

// Len returns the length of the deck (amount of cards)
func (d Deck) Len() int {
	return len(d)
}
