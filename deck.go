package main

import (
	"time"
	"math/rand"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

//Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Function to convert deck into string.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)		 
	}

	s := strings.Split(string(bs), ",")		//Convert byte to string and split the string after , and again convert it into slice
	return deck(s)
}

//Problem is that it will have last few elements fixed always because seed of rand id fixed in go everytime.
/*func (d deck) shuffle() {
	for index := range d {
		newIndex := rand.Intn(len(d) - 1)
		d[index], d[newIndex] = d[newIndex], d[i]
	}
}*/

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newIndex := r.Intn(len(d) - 1)
		d[index], d[newIndex] = d[newIndex], d[index]
	}
}