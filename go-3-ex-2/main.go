package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign)
	// TODO: Replace if, else if branching with switch/case.
	// TODO: Define all 12 cases...
	if zodiacSign == Aries {
		fmt.Println("21.03. - 20.04")
	} else if zodiacSign == Taurus {
		fmt.Println("21.04. - 21.05")
	} else {
		fmt.Println("")
	}

	switch zodiacSign {
	case 12:
		fmt.Println("21.03. - 20.04")
	case 11:
		fmt.Println("21.04. - 20.05")
	case 10:
		fmt.Println("21.05. - 20.06")
	case 9:
		fmt.Println("21.06. - 20.07")
	case 8:
		fmt.Println("21.07. - 20.08")
	case 7:
		fmt.Println("21.08. - 20.09")
	case 6:
		fmt.Println("21.09. - 20.10")
	case 5:
		fmt.Println("21.10. - 20.11")
	case 4:
		fmt.Println("21.11. - 20.12")
	case 3:
		fmt.Println("21.12. - 20.01")
	case 2:
		fmt.Println("21.01. - 20.02")
	case 1:
		fmt.Println("21.02. - 20.03")
	default:
		fmt.Println("idk")
	}

	// TODO: ...and consider a default case.
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
}
