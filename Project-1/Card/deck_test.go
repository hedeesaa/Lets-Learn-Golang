package main

import (
	"os"
	"testing"
	"time"
)

func TestNewDeck(t *testing.T) {
	firstCard := "Ace of Spades"
	lastCard := "King of Diamonds"
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 but got %v", len(d))
	}
	if d[0] != firstCard {
		t.Errorf("Expected %v but got %v", firstCard, d[0])
	}
	if d[len(d)-1] != lastCard {
		t.Errorf("Expected %v but got %v", lastCard, d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	testFile := "_deck-testing"
	os.Remove(testFile)
	deck := newDeck()
	deck.saveDeckToFile(testFile)
	time.Sleep(10 * time.Second)
	loadedDeck, err := newDeckFromFile(testFile)
	if err == nil {
		if len(loadedDeck) != 52 {
			t.Errorf("Wrote 52 to file but read %v", len(loadedDeck))
		}
	} else {
		t.Errorf("Problem on Reading the file.")
	}
	os.Remove(testFile)

}
