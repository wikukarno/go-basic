package main

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Wiku"
	middleName = "Karno"
	lastName = "Prasetyagama"

	return firstName, middleName, lastName

}

func main() {
	//you can assign the return value to multiple variables
	firstName, middleName, lastName := getFullName()
	println(firstName, middleName, lastName)

	//if you want to ignore one of the return value
	firstName, _, lastName = getFullName()
	println(firstName, lastName)
}
