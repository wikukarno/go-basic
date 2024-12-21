package main

import "fmt"

func main() {

	// If we declare a variable using "var" but do not use it,
	// it will return an error.
	// However, if we declare a constant and do not use it, it will not cause any problem.

	// const firstName string = "Wiku Karno"		// Specifying "string" is optional
	// const lastName = "Prasetyagama"
	//
	// fmt.Println(firstName)
	// fmt.Println(lastName)

	// This is how to declare multiple constants
	const (
		firstName = "Wiku Karno" // Specifying "string" is optional, same as firstName = "Wiku Karno"
		lastName  = "Prasetyagama"
	)

	// Print the constants
	fmt.Println(firstName)
	fmt.Println(lastName)
}
