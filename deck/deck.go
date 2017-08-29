package deck

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/jmuldoon/go-shuffle/card"
)

const testVersion = 0

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
	Cards  []card.Card
	length int
}

// Generate the Deck for the specified config string path,
// e.g. "config/poker.yaml"
func (d *Deck) Generate(p string) {
	cs := card.Set{}
	cs.ReadSet(p)
	d.length = len(cs.Suites) * len(cs.CardFaces)
	d.Cards = make([]card.Card, d.length)

	i := 0
	for _, s := range cs.Suites {
		for _, f := range cs.CardFaces {
			d.Cards[i] = card.Card{Suite: s, Face: f.Name, Value: f.Value}
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

// PermShuffle is a shuffle based on the use of the rand.Perm functionality
func (d *Deck) PermShuffle() {
	dst := make([]card.Card, d.length)
	perm := rand.Perm(d.length)
	for i, v := range perm {
		dst[v] = d.Cards[i]
	}
	d.Cards = dst
}

// FaroShuffle is a shuffle in which the deck is split into equal halves of 26
// cards that are then interwoven perfectly
func (d *Deck) FaroShuffle() {
	tmp := make([]card.Card, d.length)
	for i := 0; i < d.length/2; i++ {
		tmp[i*2] = d.Cards[i]
		tmp[i*2+1] = d.Cards[d.length/2+i]
	}
	d.Cards = tmp
}
