package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad "
	names[1] = "Shobir "
	names[2] = "Abdussyakur"

	fmt.Println("----------Array of String----------")
	fmt.Printf(names[0])
	fmt.Printf(names[1])
	fmt.Printf(names[2])

	var values = [3]int{
		80,
		90,
		100,
	}

	fmt.Println("\n----------Array of Number----------")
	fmt.Println(values)
	values[2] = 95
	fmt.Println(len(values))
	fmt.Println(values)
}
