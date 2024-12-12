package main

import (
	"fmt"
	"math"
)

func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := computeDiscriminant(a, b, c)
	var solutions []float64

	if D > 0 {
		solution1 := (-b + math.Sqrt(D)) / (2 * a)
		solution2 := (-b - math.Sqrt(D)) / (2 * a)
		solutions = append(solutions, solution1, solution2)
	} else if D == 0 {
		solution := -b / (2 * a)
		solutions = append(solutions, solution)
	} else {
		// Keine Lösung
	}
	return solutions
}

func main() {
	a, b, c := 3.0, 4.0, 1.0
	fmt.Printf("Lösungen (computeQuadraticFormula) mit D > 0: %v\n", computeQuadraticFormula(a, b, c)) // Output: [0.33 3.67]

	a, b, c = 2.0, 4.0, 2.0
	fmt.Printf("Lösungen (computeQuadraticFormula) mit D = 0: %v\n", computeQuadraticFormula(a, b, c)) // Output: [-1]

	a, b, c = 3.0, 4.0, 2.0
	fmt.Printf("Lösungen (computeQuadraticFormula) mit D < 0: %v\n", computeQuadraticFormula(a, b, c)) // Output: []
}
