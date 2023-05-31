package main

import "fmt"

func main() {
	names := [...]string{"Muhammad", "Shobir", "Abdussyakur", "Sinan", "Abdul", "Syakur"}
	slice := names[2:6]

	fmt.Println("1)------------ Operasi slice basic terkait index, capacity, & length ----------")
	fmt.Println("  -", slice[0])
	fmt.Println("  -", slice[1])
	fmt.Println("  -", cap(slice))
	fmt.Println("  -", len(slice))

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumaat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println("2)------------ Array sebelum diubah lewat slice -------------------------------")
	fmt.Println("  -", days)
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println("3)------------ Array setelah diubah lewat slice ke-1 --------------------------")
	fmt.Println("  -", days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println("4)------------ Array setelah diubah lewat slice ke-2 --------------------------")
	fmt.Println("  -", "Array baru:", daySlice2)
	fmt.Println("   ", "*Data diatas menjadi array baru, karena kapasitas telah dilebihi oleh input \n     datanya.")
	fmt.Println("  -", days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Muhammad"
	newSlice[1] = "Shobir"

	fmt.Println("5)------------ Membuat slice dengan command make -----------------------------")
	fmt.Println("  -", newSlice)
	fmt.Println("  -", len(newSlice))
	fmt.Println("  -", cap(newSlice))

	fmt.Println("6)------------ Mengcopy slice dengan command copy ----------------------------")
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println("  -", toSlice)
	fmt.Println("7)------------ Perbedaan array & slice ---------------------------------------")

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("  -", iniArray)
	fmt.Println("  -", iniSlice)
}
