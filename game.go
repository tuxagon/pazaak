package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/satori/go.uuid"
)

type Card struct {
	id  string
	val int
}

type Deck []Card

func CreateMainDeck() Deck {
	var deck Deck
	deck.Reset()
	deck.Shuffle()
	return deck
}

func (d *Deck) Reset() {
	var cards []Card

	func() {
		if cards == nil {
			cards = make([]Card, 40)
			for i := 0; i < len(cards); i++ {
				cards[i] = Card{
					id:  uuid.NewV4().String(),
					val: i%10 + 1}
			}
		}

		*d = Deck(cards)
	}()
}

func (d *Deck) Shuffle() {
	if *d == nil {
		d.Reset()
	}

	for i := len(*d) - 1; i > 0; i-- {
		max := i + 1
		n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		handleError(err)
		j := n.Int64()
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func handleError(e error) {
	if e != nil {
		fmt.Printf("Error: %v\n", e)
	}
}

func displayDeck(deck []Card) {
	for i := 0; i < len(deck); i++ {
		fmt.Printf("[%d]: %s\n", deck[i].val, deck[i].id)
	}
}

func main() {
	mainDeck := CreateMainDeck()
	displayDeck(mainDeck)
}
