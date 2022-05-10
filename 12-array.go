package main

import "fmt"

func main() {
	//* membuat array manual
	var names [3]string

	names[0] = "Nurcholis"
	names[1] = "Aulia"
	names[2] = "Rachman"
	
	fmt.Println(names)

	//* membuat array langsung
	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)	

	//* function array
	//? len(array) -> digunakan untuk mendapatkan panjang array, bukan jumlah data array

	fmt.Println(len(names))
	fmt.Println(len(values))
}