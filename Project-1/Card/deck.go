package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Deck Type
type deck []string

// NewDeck Generate deck of cards
func NewDeck() deck {
	cards := deck{}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardSuit := []string{"Spades", "Hearts", "Clubs", "Diamonds"}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Print Cards - Receiver
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// Deal deck
func deal(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand+1:]
}

// Shuffle deck
// It will change the deck since it is a reference type
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		randomNumber := r.Intn(len(d))
		d[index], d[randomNumber] = d[randomNumber], d[index]
	}
}

// writeToFile Write the deck to a file
func (d deck) writeToFile(address string) error {
	error := ioutil.WriteFile(address, []byte(d.toString()), 0666)
	return error
}

// toString Change deck to string
func (d deck) toString() string {
	return strings.Join(d, ", ")
}

// newDeckFromFile Returns deck from a file
func newDeckFromFile(address string) (deck, error) {
	// Read from file
	// Returns []byte
	content, err := ioutil.ReadFile(address)
	if err != nil {
		return deck{}, err
	}
	return deck(strings.Split(string(content), ",")), err

}
