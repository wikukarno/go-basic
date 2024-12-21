package main

func endApp() {
	message := recover()
	if message != nil {
		println("Error dengan message:", message)
	}
	println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	runApp(false)

}
