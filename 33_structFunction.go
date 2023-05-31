package main

import "fmt"

type Customer struct {
	Name string
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Halo", name, "My name is", customer.Name)
}

func main() {
	var admin Customer
	admin.Name = "Andi"

	admin.sayHello("joko")
}
