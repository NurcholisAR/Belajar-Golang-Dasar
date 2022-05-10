package main

import "fmt"

func main() {

	//! ketika membuat variable wajib menuliskan tipe datanya
	var name string
	name = "Nurcholis"
	fmt.Println(name)

	name = "Aulia Rachman"
	fmt.Println(name)


	//? tapi jika langsung menginisialisasi data pada variblenya maka tidak wajib menuliskan tipe datanya
	var address = "Tangerang"
	fmt.Println(address)

	var age = 20
	fmt.Println(age)

	//? membuat variable juga dapat disingkat menjadi kata kunci := 
	country := "Indonesia"
	fmt.Println(country)

	//! tetapi penggunaan := tidak menuliskan var , dan digunakan pada deklarasi awal, jika ingin merubah data tersebut maka hanya ditulis nama variable diikuti =
	country = "America"
	fmt.Println(country)

	//? digolang bisa menggunakan multiple variable
	var(
		Firstname = "Nurcholis"
		Lastname = "Rachman"
	)
	fmt.Println(Firstname)
	fmt.Println(Lastname)


}