package main

import "fmt"

func main() {
	name := "andreas"

	if name == "Muhammad" {
		fmt.Println("Halo", name)
	} else if name == "andreas" {
		fmt.Println("Halo", name)
	} else {
		fmt.Println("Kamu bukan member")
	}
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
