package main

func main() {
	// card := newCard() // do not use the colon equals when reassigning variable.
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
