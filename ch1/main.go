package main

import "fmt"

func main() {
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))
	// 	cards := deck{"Ace of Diamonds", newCard()}
	// 	cards = append(cards, "Six of Spades")
	// 	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
