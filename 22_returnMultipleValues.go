package main

import "fmt"

func getFullName() (string, string) {
	return "Muhammad", "Shobir"
}

func main() {
	// fisrtname, lastname := getFullName()
	// fmt.Println(fisrtname, lastname)
	fisrtname, _ := getFullName()
	fmt.Println(fisrtname)
}
