package main

import "fmt"

/**
Komentar
multi line
tidak terbatas
*/

func main() {
	// ini komentar single line
	sayHello("Joko")
}

func sayHello(name string) {
	fmt.Println("Program ini membahas komentar", name)
}
