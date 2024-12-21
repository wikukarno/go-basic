package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John Doe",
		"address": "New York",
	}

	// Add new key-value pair
	person["phone"] = "1234567890"

	// Update value
	person["address"] = "California"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Golang Programming"
	book["author"] = "John Doe"
	book["publisher"] = "Tech Books"
	book["price"] = "20.00"

	// Delete key-value pair
	delete(book, "price")

	fmt.Println(book)
}
