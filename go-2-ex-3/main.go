package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := make(map[int]string)
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[114] = "Modul 114:"
	// TODO: replace one
	modules[2] == "Modul 431:", modules[431]
}
