package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func NewPazaakGame() *PazaakGame {
	debug("Setting up Pazaak game...")
	game := PazaakGame{MainDeck: NewMainDeck(), MaxPlayers: 2}
	return &game
}

type PazaakGame struct {
	MainDeck      *Deck
	Players       []*Player
	CurrentPlayer *Player
	MaxPlayers    int
}

func (g *PazaakGame) Start() {
	determineFirstPlayer()
	for {
		cmd, err := fmt.Scanln("")
	}
}

func (g *PazaakGame) AddPlayer(p *Player) bool {
	if len(g.Players) == g.MaxPlayers {
		debug("No new players are permitted into this game")
		return false
	}
	debug(fmt.Sprintf("%s has joined\n", p.GamerTag))
	g.Players = append(g.Players, p)
	return true
}

func (g *PazaakGame) determineFirstPlayer() {
	debug("Determining first player...")
	n, err := rand.Int(rand.Reader, big.NewInt(int64(g.MaxPlayers)))
	handleError(err)
	g.CurrentTurn = n
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func debug(s string) {
	fmt.Println(s)
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}
