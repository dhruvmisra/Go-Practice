package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5);

	fmt.Println("hand:", hand.toString());
	fmt.Println("remainingCards:", remainingCards.toString());

	cards.saveToFile("my_cards");

	newCards := newDeckFromFile("my_cards");
	newCards.print();

	newCards2 := newDeckFromFile("my");
	newCards2.print();
}