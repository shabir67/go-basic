package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 0, "input yout number")

	flag.Parse()

	fmt.Println("Host		:", *host)
	fmt.Println("Username	:", *username)
	fmt.Println("Password	:", *password)
	fmt.Println("Password	:", *number)

}
