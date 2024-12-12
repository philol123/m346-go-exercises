package main

import (
	"errors"
	"fmt"
)

func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	if gotPoints < 0 || maxPoints <= 0 || gotPoints > maxPoints {
		return 0, errors.New("ungültige Punktzahlen: überprüfe die Eingabewerte")
	}

	grade := (gotPoints / maxPoints * 5) + 1

	return grade, nil
}

func main() {
	grade, err := computeGrade(17.5, 28.0)
	if err != nil {
		fmt.Println("Fehler:", err)
	} else {
		fmt.Println("Note für 17.5 von 28.0:", grade) // 4.125
	}

	grade, err = computeGrade(50, 50)
	if err != nil {
		fmt.Println("Fehler:", err)
	} else {
		fmt.Println("Note für 50 von 50:", grade) // 6.0
	}

	grade, err = computeGrade(30, 20)
	if err != nil {
		fmt.Println("Fehler:", err)
	} else {
		fmt.Println("Note für 30 von 20:", grade) // Fehler
	}
}
