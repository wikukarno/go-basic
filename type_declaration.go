package main

import "fmt"

func main() {
	type NoKtp string // This declares a new custom type named "NoKtp" with an underlying type of string.

	var ktpWiku NoKtp = "1111111111" // This is a variable of the custom type "NoKtp" with the value "1111111111".

	var contoh string = "2222222222"    // This is a standard string variable declaration.
	var contohKtp NoKtp = NoKtp(contoh) // This converts the string variable "contoh" into the custom type "NoKtp".

	fmt.Println(ktpWiku)   // Prints the value of the "ktpWiku" variable.
	fmt.Println(contohKtp) // Prints the value of the "contohKtp" variable, which is of type "NoKtp".
}
