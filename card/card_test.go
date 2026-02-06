package card

import (
	"cards/rank"
	"cards/suit"
	"fmt"
)

func ExampleCard_Front() {
	c1 := Card{Rank: rank.Ace, Suit: suit.Spades}
	fmt.Println(c1.Front())
	c2 := Card{Rank: rank.Ten, Suit: suit.Hearts}
	fmt.Println(c2.Front())

	// Output:
	// ┌─────┐
	// │A    │
	// │  ♠  │
	// │    A│
	// └─────┘
	// ┌─────┐
	// │10   │
	// │  ♥  │
	// │   10│
	// └─────┘
}
