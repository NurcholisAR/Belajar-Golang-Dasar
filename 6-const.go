package main

import "fmt"

func main() {
	//? membuat const wajib langsung menginisialisasi data
	const firstName string = "Nurcholis"

	//? tetapi bisa juga tidak menuliskan tipe datanya
	const lastName = "Rachman"
	const age = 20

	//? di golang jika const tidak dipanggil maka tidak terdapat error seperti var yang harus dipanggil
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	//! value conts tidak bisa diubah
	//* error
	//firstName = "Aulia"

	//? deklarasi multiple const
	const(
		city = "Tangerang"
		country = "Indonesia"
	)

	fmt.Println(city)
	fmt.Println(country)
}