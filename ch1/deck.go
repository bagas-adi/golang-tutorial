package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of deck
// which is a slice of strings

type deck []string

func main() {
	cards := newDeckFromFile("list-deck.txt")

	fmt.Println(cards.toString())

	// myHand1, remainingCards := deal(cards, 5)
	// fmt.Println("Remaining Cards : ")
	// remainingCards.print()
	// fmt.Println("My Hand : ")
	// myHand1.print()
}

func newCards() deck {
	cards := deck{}
	var cardStr string
	cardSuits := []string{"Wajik", "Hati", "Keriting", "Sekop"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "King", "Queen", "Knight", "Joker"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cardStr = cardValue + " " + cardSuit
			cards = append(cards, cardStr)
		}
	}
	// Another For loops
	// for i := range cardSuits {
	// 	for j := range cardValues {
	// 		cardStr = cardValues[j] + " " + cardSuits[i]
	// 		cards = append(cards, cardStr)
	// 	}
	// }
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func newDeckFromFile(filename string) deck {
	byteIinput, _ := ioutil.ReadFile(filename)
	strInput := string(byteIinput)
	// log.Print(err)
	return deck(strings.Split(strInput, ","))
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	var strReturn string
	for i, card := range d {
		if i == 0 {
			strReturn += card
		} else if i == len(card)-1 {
			strReturn += card
		} else {
			strReturn += ", " + card
		}
	}
	return strReturn
}

// Receiver Function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
