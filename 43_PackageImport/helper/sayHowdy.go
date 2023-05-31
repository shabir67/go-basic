package helper

import "fmt"

// Variable ataupun fungsi dengan awalan huruf kecil hanya akan berlaku di lokal.
// Jika menggunakan Huruf besar akan bisa di akses oleh file lain
func SayHowdy(name string) {
	fmt.Println("Hello", name)
}

func SayMantap(name string) {
	fmt.Println("Mantap", name)
}
