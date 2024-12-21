package main

import "fmt"

// Slice is a data structure that is used to store a collection of elements.
// A slice is a reference to an underlying array.
// A slice is a variable-length sequence which stores elements of a similar type.
// A slice is a reference type, which means that it points to the underlying array.
// A slice is a lightweight data structure that gives access to a subsequence of the elements of an array.
// A slice is a flexible and powerful tool to work with collections of elements.

// type data slice have 3 fields:
// 1. Pointer: The pointer to the first element of the array.
// 2. Length: The number of elements in the slice.
// 3. Capacity: The maximum number of elements that the slice can hold.

func main() {
	//names := []string{"Wiku", "Karno", "Prasetyagama", "Ayunda", "Putri", "Sari"}
	//slice := names[1:4]
	//
	//// Print the full name by concatenating the variables
	//println(slice[0] + " " + slice[1] + " " + slice[2])
	//
	//// Print each variable separately
	//println(slice[0])
	//println(slice[1])

	// Append slice
	//days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	//
	//daySlice1 := days[5:]
	//fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu Minggu]
	//
	//daySlice1[0] = "Sabtu Baru"
	//daySlice1[1] = "Minggu Baru"
	//fmt.Println(daySlice1) // [Sabtu Baru Minggu Baru]
	//fmt.Println(days)      // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]
	//
	//daySlice2 := append(daySlice1, "Libur Baru")
	//daySlice2[0] = "Sabtu lama"
	//fmt.Println(daySlice1)
	//fmt.Println(daySlice2)
	//fmt.Println(days)

	// Make slice
	//newSlice := make([]string, 2, 5)
	//newSlice[0] = "Wiku"
	//newSlice[1] = "Karno"
	////newSlice[2] = "Prasetyagama" // Error: Index out of range
	//
	//fmt.Println(newSlice)
	//fmt.Println(len(newSlice))
	//fmt.Println(cap(newSlice))
	//
	//newSlice2 := append(newSlice, "Prasetyagama")
	//fmt.Println(newSlice2)
	//fmt.Println(len(newSlice2))
	//fmt.Println(cap(newSlice2))
	//
	//// Copy slice
	//fromSlice := days[:]
	//toSlice := make([]string, len(fromSlice), cap(fromSlice))
	//
	//copy(toSlice, fromSlice)
	//fmt.Println(toSlice)
	//fmt.Println(toSlice)

	// difference between array and slice
	iniArray := [...]int{1, 2, 3} // array
	iniSlice := []int{1, 2, 3}    // slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
