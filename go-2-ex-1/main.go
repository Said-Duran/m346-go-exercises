package main

import "fmt"

type FullName struct {
	// TODO: add fields
}

// TODO: declare a structure for birth date

type Profile struct {
	// TODO: embed full name and birth date information
	FirstName        string
	LastName         string
	NumberOfSiblings byte
	ZodiacSign       string
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FirstName:        "Said",
		LastName:         "Duran",
		NumberOfSiblings: 4,       // TODO: adjust
		ZodiacSign:       "Libra", // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
