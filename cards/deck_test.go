package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // t is our test handler

	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v", len(d))
		// %v will place len(d) after "got"
	}

	if d[0] != "Ace of Spades" { // check 1st card
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" { // check 1st card
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFIle("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v", len(loadedDeck))
	}

	// cleanup
	os.Remove("_decktesting")
}
