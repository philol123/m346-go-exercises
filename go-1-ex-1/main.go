package main

import "fmt"

func main() {

	firstName := "Phil"
	lastName := "Hägeli"
	dayOfBirth := 10
	monthOfBirth := 03
	yearOfBirth := 2008
	numberOfSiblings := 2
	heightInMeters := 1.81
	zodiacSign := '\u2653'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
