package main

import "fmt"

func main() {
	type NoKTP string

	var KtpShobir NoKTP = "123124432432"
	fmt.Println(KtpShobir)
	fmt.Println(NoKTP("222222222222"))
}
