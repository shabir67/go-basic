package database

import "fmt"

var connection string

func init() {
	connection = "MYSQL"
	fmt.Println("Init dipanggil")
}

func GetDatabase() string {
	return connection
}
