package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	var wiku Customer
	wiku.Name = "Wiku"
	wiku.Address = "Indonesia"
	wiku.Age = 22

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     22,
	}

	fmt.Println(wiku)
	fmt.Println(joko)
	fmt.Println(Customer{"Budi", "Indonesia", 22})
	fmt.Println(Customer{Name: "Budi", Address: "Indonesia", Age: 22})

	wiku.sayHi()
}
