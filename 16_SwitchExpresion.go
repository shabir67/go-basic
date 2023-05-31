package main

import "fmt"

func main() {

	name := "Shobir"

	switch name {
	case "Shobir":
		fmt.Println("Hello", name)
	case "Joko":
		fmt.Println("Hello", name)
	default:
		fmt.Println("Halo boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama sudah lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
