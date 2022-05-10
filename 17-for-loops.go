package main

import "fmt"

func main() {

	number := 1
	for number <= 10 {
		fmt.Println("Perulangan ke", number)
		number++
	}

	//* for dengan statement
	for numbers := 10; numbers > 0; numbers-- {
		fmt.Println("perulangan ke", numbers)
	}

	//* for range
	array := [...]string{
		"Nurcholis",
		"Aulia",
		"Rachman",
	}
	for index, name := range array {
		fmt.Println("index", index, "=", name)
	}

	slice := []string{
		"Nurcholis",
		"Aulia",
		"Rachman",
	}
	for index, names := range slice {
		fmt.Println("index", index, "=", names)
	}

	address := make(map[string]string)
	address["City"] = "Tangerang"
	address["Country"] = "Indonesia"

	for key, value := range address {
		fmt.Println("Key ", key, "=", value)
	}

	//* ignore key / index jika tidak digunakan maka di tulis _

	for _, value := range address {
		fmt.Println(value)
	}
}
