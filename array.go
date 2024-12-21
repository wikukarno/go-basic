package main

import "fmt"

func main() {
	//var names [3]string
	//
	//names[0] = "Wiku"
	//names[1] = "Karno"
	//names[2] = "Prasetyagama"
	//
	//// Print the full name by concatenating the variables
	//println(names[0] + " " + names[1] + " " + names[2])
	//
	//// Print each variable separately
	//println(names[0])
	//println(names[1])
	//println(names[2])

	//// Shorthand variable declaration
	//names := [3]string{"Wiku", "Karno", "Prasetyagama"}
	//
	//// Print the full name by concatenating the variables
	//println(names[0] + " " + names[1] + " " + names[2])

	//var values = [3]int{1, 2, 3}
	//
	//// Print the full name by concatenating the variables
	//println(values[0] + values[1] + values[2])
	//
	//// Print each variable separately
	//println(values[0])
	//println(values[1])
	//println(values[2])

	var values = [...]int{1, 2, 3, 4, 5}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)

	//note
	//can't delete element in array
}
