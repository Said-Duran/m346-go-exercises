package main

import "fmt"

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	// TODO: Loop over suits and ranks to output all combinations.

	for _, rank := range suits {
		for _, suit := range ranks {
			fmt.Printf("%c%c\t", suit, rank)

		}
		fmt.Println()
	}

	cards := map[rune][]rune{
		Diamonds: {Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace},
		Spades:   {Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace},
		Clubs:    {Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace},
		Hearts:   {Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace},
	}

	fmt.Println("\nAusgabe")
	for y := 0; y < len(cards[Diamonds]); y++ {
		for suit, ranks := range cards {
			// Ausgabe
			card := fmt.Sprintf("%c%c", suit, ranks[y])
			fmt.Print(card + "\t")
		}
		fmt.Println()
	}
}
