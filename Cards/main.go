package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 8)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining Deck:")
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
