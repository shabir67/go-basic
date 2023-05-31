package main

import "fmt"

func getHello(name string) string {

	if name == "" {
		return "Hallo bosku"
	} else {
		return "Hallo " + name
	}
}

func main() {
	result := getHello("Joko")
	fmt.Println(result)
	fmt.Println(getHello(""))
	fmt.Println(getHello("budi"))

}
