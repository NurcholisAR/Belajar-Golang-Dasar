package main

import "fmt"

//* struct adalah sebuah template yg digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
//* jika di oop anggap saja seperti membuat class

//* struct adalah template data / prototype data
//* struct tidak bisa digunakan secara langsung
//* tetapi kita bisa membuat data/object dari struck yg telah kita buat
type Person struct {
	Name, Address string
	Age           int
	Status        bool
}

func main() {
	//? mengisi data struct
	var Budi Person
	Budi.Name = "Budi"
	Budi.Address = "Tangerang"
	Budi.Age = 25
	Budi.Status = false

	fmt.Println(Budi)
	fmt.Println(Budi.Name)
	fmt.Println(Budi.Address)
	fmt.Println(Budi.Age)
	fmt.Println(Budi.Status)

	//? mengisi data struct literals
	Joko := Person{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     22,
		Status:  true,
	}
	fmt.Println(Joko)

	//! tidak disarankan membuat data struct seperti dibawah, karena rentan error
	Sawi := Person{"Sawi", "Jakarta", 35, true}
	fmt.Println(Sawi)

}
