package main

import (
	"fmt"
)

const (
	Add1 = 1 << iota
	Add2
	Add3
	Add4
	Add5
	Add6
	Sub1
	Sub2
	Sub3
	Sub4
	Sub5
	Sub6
	AddOrSub1
	AddOrSub2
	AddOrSub3
	AddOrSub4
	AddOrSub5
	AddOrSub6
	AddOrSub1OrTieWin
	DoubleLastCard
	ChangeSign2Or4
	ChangeSign3Or6
	AddOrSub1Or2
)

func main() {
	mainDeck := NewMainDeck()
	fmt.Print(mainDeck.String())
	total := mainDeck.Count + 1
	for i := 0; i < total; i++ {
		card := mainDeck.DrawCard()
		if card != nil {
			fmt.Println(card.String())
		} else {
			fmt.Println("No cards left!")
		}
	}
}
