package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	fmt.Fprintln(os.Stdout, "The dice shows", eyes, "eyes")
	fmt.Fprintln(os.Stdout, "The dice was rolled at", when)

	eyesFile, _ := os.Create("eyes.txt")
	defer eyesFile.Close()
	fmt.Fprintln(eyesFile, "The dice shows", eyes, "eyes")

	logFile, _ := os.Create("dice.log")
	defer logFile.Close()
	fmt.Fprintln(logFile, "The dice was rolled at", when)
}

// Programm starten:
// go run go-1-ex-3/main.go
// Ergebnis:
// - Die gewürfelte Augenzahl wird in "eyes.txt" gespeichert.
// - Der Zeitpunkt des Würfelns wird in "dice.log" gespeichert.
