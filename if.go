package main

func main() {
	name := "Wiku"

	if name == "Wiku" {
		println("Hello Wiku")
	} else if name == "Karno" {
		println("Hello Karno")
	} else {
		println("Who are you?")
	}

	// usually use if statement
	length := len(name)

	if length > 5 {
		println("Too long")
	} else {
		println("Not too long")
	}

	//if short statement
	if length := len(name); length > 5 {
		println("Too long")
	} else {
		println("Not too long")
	}

}
