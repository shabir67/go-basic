package main

import (
	"errors"
	"fmt"
)

func Division(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Cannot Divide by Zero!")
	} else {
		return value / divider, nil
	}
}

func main() {
	result, err := Division(10, 00)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Err", err.Error())
	}
}
