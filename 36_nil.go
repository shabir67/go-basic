package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

/**
Nil dapa bekerja pada tipe data seperti:
-interface
-function
-map
-slice
pointer
-channel
*/

func main() {
	person := NewMap("Shobir")
	fmt.Println(person)
}
