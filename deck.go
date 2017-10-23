package fivetwo

import (
	"bytes"
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
