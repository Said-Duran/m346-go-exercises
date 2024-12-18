package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		Vorname  string
		Nachname string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class []Student
	// TODO: declare a map of modules being attended by multiple classes
	modules := map[string]Class{
		"Modul 117": {
			{Vorname: "Peter", Nachname: "Müller"},
		},
		"Modul 346": {
			{Vorname: "Hans", Nachname: "Müller"},
		},
	}
	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
