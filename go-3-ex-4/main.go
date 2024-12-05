package main

import "fmt"

const (
	Diamonds = "♦" // Karo
	Spades   = "♠" // Pik
	Clubs    = "♣" // Kreuz
	Hearts   = "♥" // Herz

	Six   = "6"
	Seven = "7"
	Eight = "8"
	Nine  = "9"
	Ten   = "10"
	Jack  = "J"
	Queen = "Q"
	King  = "K"
	Ace   = "A"
)

func main() {
	suits := []string{Diamonds, Spades, Clubs, Hearts}
	ranks := []string{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	fmt.Println("Karten für jedes Symbol (Farbe) und Wert:")
	for _, rank := range ranks {
		for _, suit := range suits {
			fmt.Printf("%s%s\t", suit, rank)
		}
		fmt.Println()
	}

	cards := map[string][]string{
		Diamonds: {"6", "7", "8", "9", "10", "J", "Q", "K", "A"},
		Spades:   {"6", "7", "8", "9", "10", "J", "Q", "K", "A"},
		Clubs:    {"6", "7", "8", "9", "10", "J", "Q", "K", "A"},
		Hearts:   {"6", "7", "8", "9", "10", "J", "Q", "K", "A"},
	}

	fmt.Println("\nAusgabe der Karten nach Farben:")
	for _, suit := range suits {
		fmt.Printf("\nFarbe: %s\n", suit)
		for _, card := range cards[suit] {
			fmt.Printf("%s%s\t", suit, card)
		}
		fmt.Println()
	}
}
