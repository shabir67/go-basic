package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("ma([a-z])tap")

	fmt.Println(regex.MatchString("sateminang"))
	fmt.Println(regex.MatchString("satemInang"))
	fmt.Println(regex.MatchString("satepadnang"))

	search := regex.FindAllString("mantap mabtap maztap mantab", -1)
	fmt.Println(search)
}
