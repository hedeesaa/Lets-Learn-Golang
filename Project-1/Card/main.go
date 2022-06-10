package main

import "fmt"

func main() {
	fileAddress := "cards.txt"

	cards := newDeck()
	cards.print()

	cards.saveDeckToFile(fileAddress)

	cards, _ = newDeckFromFile(fileAddress)

	cards.shuffle()

	hand, remain := deal(cards, 7)

	fmt.Println("\nThis is hand:")
	hand.print()
	fmt.Println("\nThis is remain:")
	remain.print()

}
