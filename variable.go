package main

import "fmt"

func main() {
	// Initialization with the type string

	// var name string
	//
	// name = "Wiku Karno"
	// fmt.Println(name)
	//
	// name = "Wiku Prasetyagama"
	// fmt.Println(name)

	// Continuous initialization without explicitly declaring the type string

	// var name = "Wiku Karno"		// This is the same as: var name string = "Wiku Karno"
	// fmt.Println(name)
	//
	// name = "Wiku Prasetyagama"
	// fmt.Println(name)

	// Using shorthand variable declaration (:=) without var
	// name := "Wiku Karno"
	// fmt.Println(name)
	//
	// For the next reassignment, use only the variable name (do not use :=)
	// name = "Wiku Prasetyagama"
	// fmt.Println(name)

	// Declaring multiple variables at once
	var (
		firstName  = "Wiku"
		middleName = "Karno"
		lastName   = "Prasetyagama"
	)

	// Printing the full name by concatenating the variables
	fmt.Println(firstName + " " + middleName + " " + lastName)

	// Printing each variable separately
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
