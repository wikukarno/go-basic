package main

import "fmt"

// use type declaration to define a function as a type
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

//func sayHelloWithFilter(name string, filter func(string) string) {
//	filteredName := filter(name)
//	fmt.Println("Hello", filteredName)
//}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Wiku", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
