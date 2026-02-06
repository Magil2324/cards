package card

import (
	"cards/rank"
	"cards/suit"
	"fmt"
	"strings"
)

// Card repräsentiert eine Spielkarte mit einem Wert und einer Farbe.
type Card struct {
	Rank rank.Rank
	Suit suit.Suit
}

// Front gibt die Karte als String zurück.
// Dabei soll die Karte in AsciiArt mit
// einem Rahmen dargestellt werden, z.B.:
// ┌─────┐
// │A    │
// │  ♠  │
// │    A│
// └─────┘
func (c Card) Front() string {
	rank := c.Rank
	suit := c.Suit

	lines := []string{
		"┌─────┐",
		fmt.Sprintf("│%-2s   │", rank),
		fmt.Sprintf("│  %s  │", suit),
		fmt.Sprintf("│   %2s│", rank),
		"└─────┘",
	}
	return strings.Join(lines, "\n")
}
