package structs

type Card struct {
	Name  string
	Value int
	Suit  Suit
}

type Suit string

const (
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
	Spades   Suit = "Spades"
)
