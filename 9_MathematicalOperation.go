package main

import "fmt"

func main() {
	var phi = 3.14
	var r = 7

	var lingkaran = phi * float64(r) * float64(r)
	fmt.Println("Luas lingkaran dengan jari-jari ", r, "adalah: ", lingkaran)
	fmt.Println("Hasil Augmented Assignments & Unary dari ", r)

	r *= r

	fmt.Println("Augmented ", r)

	r++

	fmt.Println("Unary(+1): ", r)

}
