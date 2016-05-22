package main

import (
	"crypto/rand"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"math/big"
	"strconv"
)

type Card struct {
	Id    string
	Value string
}

func (c *Card) String() string {
	return fmt.Sprint(c.Value)
}

func NewMainDeck() *Deck {
	d := new(Deck)
	d.Reset()
	d.Shuffle()
	return d
}

func NewSideDeck() *SideDeck {
	d := new(SideDeck)
	d.Reset()
	d.Shuffle()
	return d
}

type IDeck interface {
	Reset()
	Shuffle()
	DrawCard()
}

type Deck struct {
	Cards []*Card
	Count int
}

func (d *Deck) Reset() {
	d.Count = 40
	if d.Cards == nil {
		d.Cards = newDeck("main", d.Count)
	}
}

func (d *Deck) Shuffle() {
	fmt.Println("Shuffling...")
	for i := d.Count - 1; i > 0; i-- {
		max := i + 1
		n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		handleError(err)
		j := n.Int64()
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) DrawCard() *Card {
	fmt.Printf("Drawing Card (%d)...\n", d.Count)
	if d.Count == 0 {
		return nil
	}
	d.Count--
	return d.Cards[d.Count]
}

func (d *Deck) String() string {
	s := ""
	for i := d.Count - 1; i >= 0; i-- {
		s += fmt.Sprintf("[%d]: %s\n", i+1, d.Cards[i].Value)
	}
	return s
}

type SideDeck struct {
	Deck Deck
	Size int
}

func (d *SideDeck) Reset() {
	d.Deck.Count = d.Size
	if d.Deck.Cards == nil {
		newDeck("side", d.Deck.Count)
	}
}

func (d *SideDeck) Shuffle() {
	d.Deck.Shuffle()
}

func (d *SideDeck) DrawCard() *Card {
	return d.Deck.DrawCard()
}

func (d *SideDeck) String() string {
	return d.Deck.String()
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func newDeck(name string, size int) []*Card {
	fmt.Printf("Creating %s deck...\n", name)
	cards := make([]*Card, size)
	for i := 0; i < size; i++ {
		cards[i] = &Card{
			Id:    uuid.NewV4().String(),
			Value: strconv.Itoa(i%10 + 1)}
	}
	return cards
}
