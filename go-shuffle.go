package main

import (
	"fmt"

	"github.com/jmuldoon/go-shuffle/deck"
)

func main() {
	d := deck.Deck{}

	d.Generate("config/poker.yaml")
	fmt.Println("---- Before Sort ---")
	d.Display()
	d.FisherYatesShuffle()
	d.PermShuffle()
	d.FaroShuffle()
	fmt.Println("---- After Sort ---")
	d.Display()
}
