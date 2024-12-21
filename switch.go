package main

import "fmt"

func main() {
	name := "Karno"

	switch name {
	case "Wiku":
		fmt.Println("Hello Wiku")
	case "Karno":
		fmt.Println("Hello Karno")
	default:
		fmt.Println("Who are you?")
	}

	//short statement
	switch length := len(name); length > 5 {

	case true:
		fmt.Println("Too long")
	case false:
		fmt.Println("Not too long")
	}
}
