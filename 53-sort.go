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

func (a UserSlice) Len() int {
	return len(a)
}

//? i,j -> ini adalah index
func (a UserSlice) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func (a UserSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	users := []User{
		{"Budi", 22},
		{"Pudi", 30},
		{"Dudi", 25},
		{"Mudi", 21},
		{"Rudi", 52},
	}
	//? sebelum disort
	fmt.Println(users)

	sort.Sort(UserSlice(users))
	//? sesudah disort
	fmt.Println(users)
}
