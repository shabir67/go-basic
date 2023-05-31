package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	shobir := Man{"Shobir"}
	shobir.Married()
	fmt.Println(shobir)
}
