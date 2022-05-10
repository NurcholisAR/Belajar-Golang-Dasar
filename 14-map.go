package main

import "fmt"

func main() {

	//* deklarasi map manual
	//? make(map[tipe data key] tipe data value) -> membuat map baru
	var person map[string]string = make(map[string]string)

	//* function map
	//? map[key] = value -> untuk menambah / mengubah data di map dengan key
	person["Name"] = "Nurcholis"
	person["Age"] = "20"
	fmt.Println(person)

	//? map[key] -> untuk mengambil data di map dengan key
	fmt.Println(person["Name"])
	fmt.Println(person["Age"])

	//? len(map) -> untuk mendapatkan jumlah data di map
	fmt.Println(len(person))

	//? delete(variable map, key) -> menghapus data di map dengan key
	//* sebelum dihapus
	fmt.Println(person)
	//* setelah dihapus
	delete(person, "Age")
	fmt.Println(person)

	book := map[string]string{
		"Title":  "Belajar Go-Lang",
		"Author": "Nurcholis",
	}

	book["Upss"] = "Salah"

	fmt.Println(book)

	delete(book, "Upss")
	fmt.Println(book)

}
