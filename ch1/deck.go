package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings

type deck []string

// func main() {
// 	cards := newDeckFromFile("list-deck.txt")
// 	fmt.Println(cards.shuffle().toString())
// 	// cards.saveToFile("list-deck.txt")
// 	// myHand1, remainingCards := deal(cards, 5)
// 	// fmt.Println("Remaining Cards : ")
// 	// remainingCards.print()
// 	// fmt.Println("My Hand : ")
// 	// myHand1.print()
// }

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
	byteIinput, err := ioutil.ReadFile(filename)
	// If you don't want to use another return - byteIinput, _ := ioutil.ReadFile(filename)
	strInput := string(byteIinput)

	// Option #1 - log the err and return a call to newDeck()
	// Option #2 - log the err and entirely quit the program
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	return deck(strings.Split(strInput, ","))
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	var strReturn string
	// strReturn := strings.Join(d, ",")
	for i, card := range d {
		if i == 0 {
			strReturn += card
		} else if i == len(d)-1 {
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

func (d deck) shuffle() {
	t := time.Now().UnixNano()
	source := rand.NewSource(t)
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
