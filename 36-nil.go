package main

import "fmt"

func NewMap(name string) map[string]string {
	switch name {
	case "":
		return nil
	default:
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("Budi")
	person1 := NewMap("")

	fmt.Println(person)
	fmt.Println(person1)

	//* melakukan pengecekan manual tanpa nil / null
	var person2 map[string]string = NewMap("")
	if person2["name"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person2)
	}

	//* pengecekan menggunakan nil
	var person3 map[string]string = NewMap("Budi")
	if person3 == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person3)
	}
}
