package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Bandung", "West Java", "Indonesia"}
	address2 := &address1

	address2.City = "Sumedang"

	var address3 Address = Address{"Jaktim", "DKI Jakarta", "Indonesia"}
	var address4 *Address = &address3

	*address2 = Address{"Surabaya", "East Java", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address4)

	var address5 *Address = new(Address)
	address5.City = "Jakarta"
	fmt.Println(address5)

}
