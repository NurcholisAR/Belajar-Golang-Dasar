package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

//? untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * diparameternya
func ChangeCountryToIndonesia(alamat *Alamat) {
	alamat.Country = "Indonesia"
}

func main() {
	alamat := Alamat{
		City:     "Tangerang",
		Province: "Banten",
		Country:  "",
	}
	var alamatPointer = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
