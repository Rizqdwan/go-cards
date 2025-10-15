package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string


func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for eachSuit := range cardSuits {
		for eachValue := range cardValues {
			cards = append(cards, cardValues[eachValue] + " of " + cardSuits[eachSuit])
		}
	}

	return cards
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle(){

}

func (d deck) saveToFile(){

}