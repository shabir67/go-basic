package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("suba12 @gmail.com", "uba"))
	fmt.Println(strings.Contains("suba12 @gmail.com", "sou"))

	fmt.Println(strings.Split("Muhammad Shobir Abdussyakur", " "))

	fmt.Println(strings.ToLower("Muhammad Shobir Abdussyakur"))
	fmt.Println(strings.ToUpper("Muhammad Shobir Abdussyakur"))
	fmt.Println(strings.ToTitle("muhammad shobir abdussyakur"))

	fmt.Println(strings.Trim("       muhammad shobir abdussyakur      ", " "))

	fmt.Println(strings.ReplaceAll(" muhammad Shobir abdussyakur", "Shobir", "Shabir"))

}
