package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	//address2 := address1
	//
	//address2.City = "Bandung"
	//
	//fmt.Println(address1) // not change
	//fmt.Println(address2)	// change

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1) // change
	fmt.Println(address2) // change
}
