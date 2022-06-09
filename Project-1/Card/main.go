package main

import "fmt"


func main(){
	cards := NewDeck()

	fmt.Println("This is the Shuffled Card: ")
	cards.shuffle()
	cards.print()

	fmt.Println("\nHand of 7:")
	hand1, _ := deal(cards, 7)
	hand1.print()
}