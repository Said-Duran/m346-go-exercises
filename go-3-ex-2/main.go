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

	switch zodiacSign {
	case 9800:
		fmt.Println("21.03. - 20.04")
	case 9801:
		fmt.Println("21.04. - 20.05")
	case 9802:
		fmt.Println("21.05. - 20.06")
	case 9803:
		fmt.Println("21.06. - 20.07")
	case 9804:
		fmt.Println("21.07. - 20.08")
	case 9805:
		fmt.Println("21.08. - 20.09")
	case 9806:
		fmt.Println("21.09. - 20.10")
	case 9807:
		fmt.Println("21.10. - 20.11")
	case 9808:
		fmt.Println("21.11. - 20.12")
	case 9809:
		fmt.Println("21.12. - 20.01")
	case 9810:
		fmt.Println("21.01. - 20.02")
	case 9811:
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
