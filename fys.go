package main

import (
	"fmt"

	"github.com/jmuldoon/fisher-yates-shuffle/deck"
)

func main() {
	d := deck.Deck{}

	d.Generate()
	fmt.Println("---- Before Sort ---")
	d.Display()
	d.FisherYatesShuffle()
	fmt.Println("---- After Sort ---")
	d.Display()
}
