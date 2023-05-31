package main

import (
	"fmt"

	"github.com/shabir67/Basic/pzn/Dasar/43_PackageImport/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
