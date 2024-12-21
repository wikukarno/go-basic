package main

func getSayHello(name string) string {
	hello := "hello " + name

	return hello
}

func main() {
	result := getSayHello("Wiku")
	println(result)
}
