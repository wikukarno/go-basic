package main

import "fmt"

func factorialLoop(value int) int {
	result := 1 // 1 is the neutral value for multiplication

	// loop from value to 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {

	// Factorial with recursive
	result := factorialRecursive(5)
	fmt.Println(result)

	// Factorial with loop
	result2 := 5 * 4 * 3 * 2 * 1
	fmt.Println(result2)
	result3 := factorialLoop(5)
	fmt.Println(result3)
}
