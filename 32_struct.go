package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var pelanggan Customer
	pelanggan.Name = "Shobir"
	pelanggan.Address = "Dago"
	pelanggan.Age = 30

	pelanggan2 := Customer{
		Name:    "Airlangga",
		Address: "Tuban",
		Age:     72,
	}

	pelanggan3 := Customer{"Mega", "Indo", 99}

	fmt.Println(pelanggan)
	fmt.Println(pelanggan2)
	fmt.Println(pelanggan3)
}
