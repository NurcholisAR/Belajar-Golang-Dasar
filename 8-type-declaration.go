package main

import "fmt"

func main() {
	//? membuat alias untuk tipe data di golang
	type noKTP string
	type married bool

	var ktpSaya noKTP = "123242213"
	var status married = false

	fmt.Println(ktpSaya)
	fmt.Println(status)
}