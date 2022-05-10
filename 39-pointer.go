package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//* secara default digolang semua variable itu di passing by value, bukan by reference
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	//? pass by value
	/**
	        var address2 Address = address1
	        address2.City = "Bandung"
	        fmt.Println(address1) -> address1 tidak berubah
	        fmt.Println(address2)
		**/
	//? pass by reference
	// var address2 *Address = &address1
	address2 := &address1
	address2.City = "Bandung"

	//* Operator *
	//? jika ingin mengubah seluruh variable yang mengacu pada data tersebut (cth: mengacu ke address1), kita bisa menggunakan operator *
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	var address3 *Address = &address1

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	//* Function New
	//* golang juga memiliki function new yang bisa digunakan untuk membuat pointer
	//* Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	// dilakukan secara manual
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}
	var address4 *Address = new(Address)

	address5 := address4
	address5.City = "Tangerang"
	address5.Country = "Indonesia"

	fmt.Println(address4) // -> address4 berubah
	fmt.Println(address5)
}
