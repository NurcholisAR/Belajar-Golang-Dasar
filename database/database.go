package database

import "fmt"

//* saat kita membuat package, kita bisa membuat sebuah function yg akan diakses ketika package kita diakses

//* ini sangat cocok ketika contohnya, jika package kita berii function" untuk berkomunikasi dangan databse, kita membuat function inisialisasi untuk membuka koneksi ke database

//* untuk membuat function yang diakses secara otomatis ketika package diakses, kita cuku membuat function dengan nama init

var connection string

func init() {
	fmt.Println("init dieksekusi")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
