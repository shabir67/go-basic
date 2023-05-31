package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, number := range number {
		total += number
	}
	return int(total)

}

func main() {
	totalSum := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Hasil fungsi variadic dg multi arguments: ", totalSum)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	totalSlice := sumAll(slice...)
	fmt.Println("Hasil fungsi variadic dg input slice: ", totalSlice)
}
