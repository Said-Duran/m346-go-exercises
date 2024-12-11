package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	// TODO: Implement FizzBuzz using a for loop from Lower to Upper.

	for i := 0; i < 31; i++ {
		if i%3 == 0 {
			fmt.Println(i, ": ", "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, ": ", "Buzz")
		} else if i%5 == 0 && i%3 == 0 {
			fmt.Println(i, ": ", "FizzBuzz")
		} else {
			fmt.Println(i, ": ", i)
		}
	}
}
