package main

func main() {
	counter := 1

	for counter <= 10 {
		println("Perulangan ke", counter)
		counter++
	}

	// for with statement
	for i := 0; i < 10; i++ {
		println("Perulangan ke", i)
	}

	// for with break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println("Perulangan ke", i)
	}

	// for with continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		println("Perulangan ke", i)
	}

	// for with range
	names := []string{"Wiku", "Karno", "Naruto"}
	for i, name := range names {
		println("Perulangan ke", i, "dengan nama", name)
	}

	// for with range and ignore index
	for _, name := range names {
		println("Nama", name)
	}

	// for with range and ignore value
	for i := range names {
		println("Index", i)
	}
}
