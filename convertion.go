package main

import "fmt"

func main() {
	var nilai32 int32 = 32768          // Declare a 32-bit integer variable
	var nilai64 int64 = int64(nilai32) // Convert int32 to int64
	var nilai16 int16 = int16(nilai32) // Convert int32 to int16

	fmt.Println(nilai32) // 32768: This is the original value of the int32 variable.
	fmt.Println(nilai64) // 32768: The value remains the same because int64 has enough space to hold it.
	fmt.Println(nilai16) // -32768: This value wraps around because it exceeds the maximum range of int16.

	var name = "Wiku Karno" // Declare a string variable
	var w uint8 = name[0]   // Get the byte value of the first character in the string (index 0)
	var wString = string(w) // Convert the byte value into a string

	fmt.Println(name)    // This is the original value of the variable "name".
	fmt.Println(w)       // This is the byte value of the first character in the string "name" (from index 0).
	fmt.Println(wString) // This is the string representation of the byte value stored in "w".
}
