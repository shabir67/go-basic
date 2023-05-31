package main

import "fmt"

type Person struct {
	Name string
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("My name is", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Shobir"}
	SayHello(person)

	cat := Animal{
		Name: "Puss",
	}
	SayHello(cat)
}
