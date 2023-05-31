package main

import (
	"fmt"
)

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println("String", value, "is a string")
	case int:
		fmt.Println("int", value, "is an integer")
	default:
		fmt.Println("Unknown data types")
	}
}
