package main

func NewPlayer(gamerTag string) *Player {
	p := Player{GamerTag: gamerTag, SideDeck: NewSideDeck()}
	return &p
}

type Player struct {
	GamerTag string
	SideDeck *SideDeck
}
