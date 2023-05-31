package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perluangan ke", counter)
	// 	counter = counter + 1
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perualangan ke", counter)
	// }

	// slice := [...]string{"muhammad", "shobir", "abdussyakur"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }
	// Tanpa index
	// for _, value := range slice {
	// 	fmt.Println("Index", "=", value)
	// }
	// Dengan index
	// for i, value := range slice {
	// 	fmt.Println("Index", i, "=", value)
	// }

	person := make(map[string]string)
	person["name"] = "Shobir"
	person["title"] = "Orang biasa"

	for key, value := range person {
		fmt.Println(key, ":", value)
	}

}
