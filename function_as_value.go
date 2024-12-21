package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name

}

func main() {
	contoh := getGoodBye
	result := contoh("Wiku")
	fmt.Println(result)
}
