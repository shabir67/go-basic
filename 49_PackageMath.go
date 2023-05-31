package main

import (
	"fmt"
	"math"
)

func main() {
	value1 := math.Round(2.99)
	value2 := math.Floor(2.99)
	value3 := math.Ceil(2.10)
	value4 := math.Max(2.9, 3.14)
	value5 := math.Min(2.9, 3.14)

	fmt.Println("Round	: ", value1)
	fmt.Println("Floor	: ", value2)
	fmt.Println("Ceil	: ", value3)
	fmt.Println("Max	: ", value4)
	fmt.Println("Min	: ", value5)

}
