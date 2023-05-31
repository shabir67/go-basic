package main

import "fmt"

// Gunakan pointer untuk data yang memiliki field yang banyak

type Address struct {
	City, Province, Country string
}

func ChengeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address6 = Address{
		City:     "Cimahi",
		Province: "West Java",
		Country:  "",
	}
	ChengeCountryToIndonesia(&address6)
	fmt.Println(address6)
}
