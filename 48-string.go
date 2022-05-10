package main

import (
	"fmt"
	"strings"
)

func main() {
	var Name string = "Nurcholis Aulia Rachman"

	//* Beberapa function di package strings

	//* strings.Contains(string, search)
	//? mengecek apakah string mengandung string lain
	fmt.Println(strings.Contains(Name, "Rachman"))
	fmt.Println(strings.Contains(Name, "Rahman"))

	//* strings.ToLower(string)
	//? membuat semua karakter string menjadi lower case
	fmt.Println(strings.ToLower(Name))

	//* strings.ToUpper(string)
	//? membuat semua karakter string menjadi upper case
	fmt.Println(strings.ToUpper(Name))

	//* strings.Split(string, separator)
	//? memotong string berdasarkan separator
	fmt.Println(strings.Split(Name, ""))
	fmt.Println(strings.Split(Name, " "))

	//* strings.Trim(string, cutset)
	//? memotong cutset diawal dan diakhir string
	fmt.Println(strings.Trim("          Nurcholis      ", " "))

	//* strings.ReplaceAll(string, from, to)
	//? mengubah semua string dari from ke to
	fmt.Println(strings.ReplaceAll("kopi kopi kopi", "kopi", "gula"))

}
