package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result = random()
	//var resultString = result.(string)
	//fmt.Println(resultString)

	//var resultInt = result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
