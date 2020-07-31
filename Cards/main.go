package main

import (
	"fmt"
	"os"
)

func main() {
	/*cards := newDeck()

	hand, remainingDeck := deal(cards, 8)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining Deck:")
	remainingDeck.print()

	fmt.Println(cards.toString())

	filename, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	filename += "/Cards/CardsFileTest"

	cards.saveToFile(filename)*/

	filename, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	filename += "/Cards/CardsFileTest"
	cardsFromFile := newDeckFromFile(filename)
	cardsFromFile.shuffle()
	cardsFromFile.print()
}

func newCard() string {
	return "Five of Diamonds"
}
