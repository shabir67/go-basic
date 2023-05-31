package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Muhammad")
	data.PushBack("Shobir")
	data.PushBack("Abdussyakur")
	data.PushFront("Mr")

	// fmt.Println((data.Front().Value))
	// fmt.Println((data.Back().Value))
	// fmt.Println((data.Back().Prev().Value))
	// fmt.Println(data)
	// Hasil Loop dari depan
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	// Hasil loop dari belakang
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
