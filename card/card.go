package card

// FaceEnum basic card faces
type FaceEnum int

// ValueEnum card values
type ValueEnum int

// Enumerals for the faces
const (
	CLUB FaceEnum = iota
	DIAMOND
	HEART
	SPADE
)

// Enumerals for the values
const (
	ACE ValueEnum = 1 + iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var faces = []string{
	"Club",
	"Diamond",
	"Heart",
	"Spade",
}

var values = []string{
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
}

// String returns the English name of the face (Club, Diamond, Heart, Spade).
func (f FaceEnum) String() string { return faces[f] }

// String returns the English name of the value (Ace, One, ..., Jack, Queen, King)
func (v ValueEnum) String() string { return values[v] }

// Card is the structure holding the FaceEnum and ValueEnum of each card
type Card struct {
	Face  string
	Value string
}

// Faces returns the string array for all the card faces
func Faces() []string { return faces }

// Values returns the string array for all the card values
func Values() []string { return values }
