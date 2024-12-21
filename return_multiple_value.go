package main

func getFullNames() (string, string) {
	return "Wiku", "Karno"
}

func main() {
	//you can assign the return value to multiple variables
	//firstName, lastName := getFullNames()
	//println(firstName, lastName)

	//if you want to ignore one of the return value
	firstName, _ := getFullNames()
	println(firstName)
}
