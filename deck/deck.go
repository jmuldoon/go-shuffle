package deck

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/jmuldoon/fisher-yates-shuffle/card"
)

// DeckSize is the number of cards contained in the deck
const (
	DeckSize = 52
)

// Manager is the interface by which the Deck in manipulated
type Manager interface {
	Display()
	Generate()
	Shuffle()
}

// Deck is the structure holding the Cards
type Deck struct {
	Cards []card.Card
}

// Generate the Deck
func (d *Deck) Generate() {
	d.Cards = make([]card.Card, DeckSize)

	i := 0
	for _, f := range card.Faces() {
		for _, v := range card.Values() {
			d.Cards[i] = card.Card{Face: f, Value: v}
			i++
		}
	}
}

// Display prints all the cards in the deck
func (d *Deck) Display() {
	for _, cd := range d.Cards {
		fmt.Println(cd)
	}
}

// FisherYatesShuffle is an improved version (Durstenfeld) of the original
// Fisher-Yates algorithm with O(n) time complexity
func (d *Deck) FisherYatesShuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// // for i, val := range d.Cards {
	for i := len(d.Cards) - 1; i > 0; i-- {
		index := int(math.Floor(float64(r.Intn(52))))

		// SWAP
		tmp := d.Cards[index]
		d.Cards[index] = d.Cards[i]
		d.Cards[i] = tmp
	}
}
