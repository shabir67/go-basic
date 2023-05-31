package main

import "fmt"

func loging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer loging()
	fmt.Println("Run Application")
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// runApplication()
	runApp(true)
}
