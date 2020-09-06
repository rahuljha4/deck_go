package main

// import "fmt"

func main() {
	// card := "Ace of Spades"
	// card := newCard()

	cards := newDeck()

	//For loop in go
	/*for i, card := range cards {
		fmt.Println(i, card)
	}*/

	// hand, remaningDeck := deal(cards, 5)
	// hand.print()
	// remaningDeck.print()

	//Create a file and then read it to show data.
	/*cards.saveToFile("my_deck")
	cards = newDeckFromFile("my_deck")
	cards.print()*/

	//Shuffle the card deck
	cards.shuffle()
	cards.print()
}

