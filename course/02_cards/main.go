package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 15)

	cards.print()
	fmt.Println("----------")
	hand.print()
	fmt.Println("----------")
	remainingDeck.print()
}
