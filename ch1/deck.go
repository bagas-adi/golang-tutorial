package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

type deck []string

func main() {
	cards := newCards()
	cards.print()
}

func newCards() deck {
	cards := deck{}
	var cardStr string
	cardSuits := []string{"Wajik", "Hati", "Keriting", "Sekop"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "King", "Queen", "Knight", "Joker"}

	for i := range cardSuits {
		// fmt.Println(cardSuits[i])
		for j := range cardValues {
			cardStr = cardValues[j] + " " + cardSuits[i]
			cards = append(cards, cardStr)
			// fmt.Println(cardStr)
		}
	}
	return cards
}

// Receiver Function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
