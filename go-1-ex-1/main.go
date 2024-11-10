package main

import "fmt"

func main() {
	var firstName string = "Said"
	lastName := "Duran"
	var dayOfBirth byte = 2
	monthOfBirth := 10
	yearOfBirth := 2006
	numberOfSiblings := 4
	var heightInMeters float32 = 1.75
	var zodiacSign = '\u264e'
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
