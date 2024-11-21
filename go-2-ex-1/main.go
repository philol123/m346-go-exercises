package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{
			FirstName: "Phil",
			LastName:  "Hägeli",
		},
		BirthDate: BirthDate{
			Day:   10,
			Month: 3,
			Year:  2008,
		},
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       '♓', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
