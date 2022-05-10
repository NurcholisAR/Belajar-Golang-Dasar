package main

import "fmt"

func main() {

	name := "AA"

	switch name {
	case "Nurcholis":
		fmt.Println("Hello, Nurcholis")
		fmt.Println("Hello,", name)

	case "Aulia":
		fmt.Println("Hello, Aulia")
		fmt.Println("Hello,", name)

	case "Rachman":
		fmt.Println("Hello, Rachman")
		fmt.Println("Hello,", name)

	default:
		fmt.Println("Hai, kenalan donk")
	}

	//* switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//* switch tanpa kondisi
	panjang := len(name)

	switch {
	case panjang > 10:
		fmt.Println("Nama terlalu panjang")
	case panjang > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
