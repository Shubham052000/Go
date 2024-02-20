package main

func main() {
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 15)

	// cards.print()
	// fmt.Println("----------")
	// hand.print()
	// fmt.Println("----------")
	// remainingDeck.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
