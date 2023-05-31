package main

import "fmt"

func main() {
	var name1 = "ismi"
	var name2 = "ismi"

	var result bool = name1 < name2

	fmt.Println("-----------String Comparison-----------")
	fmt.Println(result)
	fmt.Println("-----------Integer Comparison-----------")

	var value1 = 100
	var value2 = 200

	fmt.Println("Value1>Value2", value1 > value2)
	fmt.Println("Value1<Value2", value1 < value2)
	fmt.Println("Value1==Value2", value1 == value2)
	fmt.Println("Value1!=Value2", value1 != value2)
}
