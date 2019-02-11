package main

func main() {
	cards := newDeck()
	cards.shuffle()
	hand1, remainingDeck := deal(cards, 5)
	hand2, remainingDeck := deal(remainingDeck, 5)
	hand3, remainingDeck := deal(remainingDeck, 5)
	hand1.print()
	hand2.print()
	hand3.print()
}
