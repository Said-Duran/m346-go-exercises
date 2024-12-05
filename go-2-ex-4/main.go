package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type student struct {
		Vorname  string
		Nachname string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	var class = []string{}
	class = append(class, "Max Mustermann")
	class = append(class, "Lea Müller")
	class = append(class, "Peter Meier")
	// TODO: declare a map of modules being attended by multiple classes
	var modules map[int]string
	modules[0] = "Modul 114"
	modules[1] = "Modul 257"
	modules[2] = "Modul 320"
	// TODO: output everything using fmt.Println()
	fmt.Println(class)
}
