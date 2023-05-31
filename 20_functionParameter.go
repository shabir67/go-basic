package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	name := "muhammad"
	sayHelloTo(name, "shobir")
	sayHelloTo("Teddy", "Minahasa")
}
