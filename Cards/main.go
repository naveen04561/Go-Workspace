package main

import "fmt"

func main() {
	cards := newDeck()

	//cards.print()
	//hand, remainingCards := deal(cards, 3)

	// hand.print()
	// remainingCards.print()

	// byte slice
	// str := "Hello! "
	// fmt.Println([]byte(str))

	fmt.Println(cards.toString())
}
