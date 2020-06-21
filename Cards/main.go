package main

func main() {
	cards := newDeck()
	cards.saveToFile("example.txt")
	/*handSize := 6
	hand, remainingDeck := deal(cards, handSize)
	hand.print()
	remainingDeck.print()*/
	//fmt.Println(cards.toString())
	cards2 := newDeckFromFile("example.txt")
	cards2.shuffle()
	cards2.print()
}
