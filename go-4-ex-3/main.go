package main

import (
	"fmt"
	"math"
)
// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) float64 {
     d := math.Pow(b, 2) - 4 * a*c 
	 x := (-b +- math.Sqrt(math.Pow(b,2) - 4*a*c) / 2*a )
	 return d
	 return x
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(12, 18, 9))
}
