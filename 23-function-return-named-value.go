package main

import "fmt"

func getPerson() (name string, age int, status bool) {
	//* jika tipe data return sama semua bisa disingkat menjadi
	//* func getFullName() (firstName, middlename, lastName string)
	name = "Nurcholis"
	age = 22
	status = true

	return name, age, status
	//* bisa juga hanya di tulis 'return', karena tipe data variable sudah dideklarasikan diatas
}

func main() {
	//* nama variable bebas
	name, age, status := getPerson()

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(status)
}
