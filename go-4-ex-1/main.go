package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeGrade
func computeGrade(gotPoints float64, maxPoints float64) float64 {
	finalGrade := gotPoints/maxPoints*5 + 1
	return finalGrade
}

func main() {
	// TODO: call the function computeGrade
	x1 := computeGrade(19, 47)
	fmt.Println(math.Floor(x1))
	x2 := computeGrade(54, 71)
	fmt.Println(math.Floor(x2))
	x3 := computeGrade(20, 87)
	fmt.Println(math.Round(x3))
}
