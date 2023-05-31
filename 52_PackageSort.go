package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {

	users := []User{
		{"Imre", 45},
		{"Shobir", 25},
		{"Joko", 35},
		{"Mustafa", 55},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

}
