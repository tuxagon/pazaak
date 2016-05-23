package main

import (
	"fmt"
)

func main() {
	game := NewPazaakGame()
	game.AddPlayer(NewPlayer("tuxagon91"))
	game.AddPlayer(NewPlayer("tiger411"))
	game.Start()
}
