package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Expected deck length is 24, but got %d", len(d))
	}
	if d[0] != "Queen of Spades" {
		t.Errorf("Expected first card is Quuen of Spades, but got %v", d[0])
	}
}
func Test_IO(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 24 {
		t.Errorf("Expected deck length is 24, but got %d", len(d))
	}
	os.Remove("_decktesting")
}
