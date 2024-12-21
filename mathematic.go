package main

import "fmt"

func main() {
	// Arithmetic
	var a = 10
	var b = 90
	var c = a + b
	fmt.Println(c)

	// Augmented Assignment
	var i = 10
	i += 10
	fmt.Println(i) // 20

	// Comparison
	var x = 10
	var y = 10
	fmt.Println(x == y) // true

	// Logical
	var l = true
	var m = false
	fmt.Println(l && m) // false
	fmt.Println(l || m) // true
	fmt.Println(!l)     // false

	// Unary
	var n = 10
	n++
	fmt.Println(n) // 11 because n++ is equal to n = n + 1
	n--
	fmt.Println(n) // 10 because n-- is equal to n = n - 1
	n = -n
	fmt.Println(n) // -10 because -n is equal to n = -n

}
