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
	/*for _, x := range suits {
		fmt.Println(string(x))
	}
	for _, n := range ranks {
		fmt.Println(string(n))
	} */

	for x := range suits {
		for _, y := range ranks {
			fmt.Printf(string(y), string(x))
			fmt.Printf("%c%c\t", suits, ranks)

		}
	}
}
