package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of deck
//shoudld be a slice of strings
type deck []string      // new type
var totalCards int = 12 // dynamic declaration does'nt work

// functions

func deal(d deck, cards int) (deck, deck) {

	// cards represent how many cards we need to deal

	return d[:cards], d[cards:]

}

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three"}

	for i := range cardSuits {
		for j := range cardValues {
			cards = append(cards, cardValues[j]+" of "+cardSuits[i])
		}
	}

	return cards
}

func newDeckFromFile(fileName string) deck {

	// bs is our byte slice
	bs, err := ioutil.ReadFile(fileName)

	if err != nil { // if we found error
		// error handling
		// log the error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // exit the programme
	}

	// now we need to convert the byte slice -> string -> deck
	s := strings.Split(string(bs), ",")
	// now convert this slice of string to deck

	return deck(s)

}

// recievers

func (d deck) print() { // reciever function
	for i := range d {
		fmt.Println(d[i])
	}
}

func (d deck) toString() string {

	// converting deck to string with seperator as ","
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFIle(fileName string) error {
	// converting deck to byte slice
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) Shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1) // Pseudo random generator

		d[i], d[newPosition] = d[newPosition], d[i] // swapping
	}
}
