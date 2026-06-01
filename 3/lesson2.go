package main

import "fmt"

func main() {
	var favoriteColor = "violet"
	fmt.Println("My favorite color is", favoriteColor)
	// color

	birthDate, ageYears := 1999, 27
	fmt.Println("I was born in -", birthDate, "and i'm", ageYears, "years old")
	// date of birth and age

	var (
		firstNameInitial = "S"
		lastNameInitial  = "B"
	)
	fmt.Println("Initials =", firstNameInitial, lastNameInitial)
	// block assignment and initials

	var ageInDays int
	ageInDays = ageYears * 365
	fmt.Println("My age in days is approximately", ageInDays)
	// age in days
}
