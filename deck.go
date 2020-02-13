package main

import "fmt"

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	//nested for loop
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// Receiver function
// Any variable of type deck can use this
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
