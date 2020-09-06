package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 48 {
		t.Errorf("Expected deck of 48 size got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades got %v", d[0])
	}

	if d[len(d) -1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs got %v", d[len(d) -1])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting") 

	if len(loadedDeck) != 48 {
		t.Errorf("Expected deck of size 48 got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}