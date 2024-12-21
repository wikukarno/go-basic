package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	wiku := Man{"Wiku"}
	wiku.Married()

	fmt.Println(wiku.Name)
}
