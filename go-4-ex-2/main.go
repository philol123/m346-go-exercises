package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func (s ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func main() {
	a, b := 3.0, 4.0
	fmt.Printf("Länge der Hypotenuse (computeHypotenuse): %.2f\n", computeHypotenuse(a, b)) // Output: 5.00

	sides := ShortSides{a: 5.0, b: 12.0}
	fmt.Printf("Länge der Hypotenuse (Hypotenuse Methode): %.2f\n", sides.Hypotenuse()) // Output: 13.00
}
