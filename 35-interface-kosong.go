package main

import "fmt"

//* interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan mengimplementasi nya
func Ups() interface{} {
	//return 1
	// return true
	return "Ups"
}

func val(angka int) interface{} {
	switch angka {
	case 1:
		return 1
	case 2:
		return true
	default:
		return "Ups"
	}
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)

	//! harus menggunakan tipe data interface kosong
	var value interface{} = val(4)
	fmt.Println(value)

}
