package main

func main() {
	// var card string = "Ace of Spades"
	cards := deck{newCard(), "Ace of Diamonds"} 

	// Add a value
	cards = append(cards, "Six of Spades");

	cards.print();
}

func newCard() string {
	return "Five of Diamonds"
}