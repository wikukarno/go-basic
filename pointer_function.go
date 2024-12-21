package main

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func changeCountryToIndonesiaWithPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address2 := address1
	address2.Country = "Amerika Serikat"

	println(address1.Country) // Indonesia
	println(address2.Country) // Amerika Serikat

	address3 := Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}

	changeCountryToIndonesia(&address3)
	println(address3.Country) // Indonesia

	address4 := Address{
		City:     "Surabaya",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	changeCountryToIndonesiaWithPointer(&address4)
	println(address4.Country) // Indonesia
}
