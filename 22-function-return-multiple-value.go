package main

import "fmt"

//* jika ingin mengembalikan 2 data, maka wajib menambahkan () di tipe data return value
func getName() (string, string, int) {
	return "Nurcholis", "Rachman", 20
}

func main() {
	firstName, lastName, Age := getName()

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(Age)

}
