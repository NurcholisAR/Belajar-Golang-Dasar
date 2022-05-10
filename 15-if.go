package main

import "fmt"

func main() {

	name := "abc"

	if name == "Nurcholis" {
		fmt.Println("Hello ", name)
	} else if name == "abc" {
		fmt.Println("Hello ", name)
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	//* if short statement
	//* sama seperti
	//? var length = len(name)
	//? if length < 5
	if length := len(name); length < 5 {
		fmt.Println("Nama sudah benar")
	} else {
		fmt.Println("Nama terlalu panjang")
	}
}
