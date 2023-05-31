package main

import "fmt"

func getCompleteName() (firstName, midName, lastName string) {
	firstName = "Muhammad"
	midName = "Shobir"
	lastName = "Abdussyakur"
	return
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
