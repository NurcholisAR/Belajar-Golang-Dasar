package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)

	//? jika mengkonversi yg lebih kecil, maka terjadi integer overflow. jadi saat dia sudah mencapai batas maka akan mengulang dari paling minimum (int8 : minimum = -128, maksimum = 127) dan akan terus naik dari paling minimum sampai maksimum 
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//* menkonversi ke string
	name := "Nurcholis"
	e := name[0]
	eString := string(e)

	fmt.Println(name)
	fmt.Println(eString)

	
}