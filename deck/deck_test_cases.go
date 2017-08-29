package deck

import (
	"fmt"

	"github.com/jmuldoon/go-shuffle/card"
)

var errDefault = fmt.Errorf("deck: expected error")

type Expected struct {
	Value interface{}
	Error error
}

type Tested struct {
	Value interface{}
	Error error
}

var pokerDeck = []struct {
	Suite string
	Face  string
	Value int
}{
	{"Club", "Two", 2},
	{"Club", "Three", 3},
	{"Club", "Four", 4},
	{"Club", "Five", 5},
	{"Club", "Six", 6},
	{"Club", "Seven", 7},
	{"Club", "Eight", 8},
	{"Club", "Nine", 9},
	{"Club", "Ten", 10},
	{"Club", "Jack", 11},
	{"Club", "Queen", 12},
	{"Club", "King", 13},
	{"Club", "Ace", 14},
	{"Spade", "Two", 2},
	{"Spade", "Three", 3},
	{"Spade", "Four", 4},
	{"Spade", "Five", 5},
	{"Spade", "Six", 6},
	{"Spade", "Seven", 7},
	{"Spade", "Eight", 8},
	{"Spade", "Nine", 9},
	{"Spade", "Ten", 10},
	{"Spade", "Jack", 11},
	{"Spade", "Queen", 12},
	{"Spade", "King", 13},
	{"Spade", "Ace", 14},
	{"Heart", "Two", 2},
	{"Heart", "Three", 3},
	{"Heart", "Four", 4},
	{"Heart", "Five", 5},
	{"Heart", "Six", 6},
	{"Heart", "Seven", 7},
	{"Heart", "Eight", 8},
	{"Heart", "Nine", 9},
	{"Heart", "Ten", 10},
	{"Heart", "Jack", 11},
	{"Heart", "Queen", 12},
	{"Heart", "King", 13},
	{"Heart", "Ace", 14},
	{"Diamond", "Two", 2},
	{"Diamond", "Three", 3},
	{"Diamond", "Four", 4},
	{"Diamond", "Five", 5},
	{"Diamond", "Six", 6},
	{"Diamond", "Seven", 7},
	{"Diamond", "Eight", 8},
	{"Diamond", "Nine", 9},
	{"Diamond", "Ten", 10},
	{"Diamond", "Jack", 11},
	{"Diamond", "Queen", 12},
	{"Diamond", "King", 13},
	{"Diamond", "Ace", 14},
}

func generateTestDeck() []card.Card {
	var cards = []card.Card{}
	for _, c := range pokerDeck {
		cards = append(cards, c)
	}
	return cards
}

var testDeck = generateTestDeck()

var testPokerDeckCases = []struct {
	Description string
	Expected
	Tested
}{
	{"STD Poker Deck Unshuffled",
		Expected{
			Value: &Deck{
				Cards:  testDeck,
				length: 52,
			},
			Error: error(nil),
		},
		Tested{
			Value: "../config/poker.yaml",
			Error: error(nil),
		},
	},
}
