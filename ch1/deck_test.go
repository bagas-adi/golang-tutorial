package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newCards()
	if len(d) != 52 {
		t.Errorf("Expected deck lenth of 52, but got %v", len(d))
	}
	if d[0] != "Ace Wajik" {
		t.Errorf("Expected deck first card of Ace Wajik but got %v", d[0])
	}
	if d[len(d)-1] != "Joker Sekop" {
		t.Errorf("Expected deck first card of Joker Sekop but got %v", d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T){ 
	d := newCards()
	os.Remove("_decktesting")
	err = d.saveToFile("_decktesting")
	if (err)
	os.IsExist("_decktesting")
}