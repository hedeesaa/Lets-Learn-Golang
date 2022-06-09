package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck Type
type deck []string


// Generate deck of cards
func NewDeck() deck {
	cards := deck{}
	cardValue := []string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}
	cardSuit := []string{"Spades","Hearts","Clubs","Diamonds"}

	for _, suit := range cardSuit{
		for _, value := range cardValue{
			cards = append(cards,value +" of "+suit)
		}
	}
	return cards
}


// Print Cards - Reciever
func (d deck) print(){
	for index, card := range d{
		fmt.Println(index, card)
	}
}

// Deal deck
func deal(d deck, hand int) (deck, deck){
	return d[:hand], d[hand+1:]
}

// Shuffle deck
// It will change the deck since it is a reference type
func (d deck) shuffle (){
	source := rand.NewSource(time.Now().UnixNano())
	r :=rand.New(source)
	for index := range d{
		randomNumer := r.Intn(len(d))
		d[index],d[randomNumer] = d[randomNumer],d[index]
	}
}

// TODO: Create the deck from the file

// TODO: Write the deck to a file


