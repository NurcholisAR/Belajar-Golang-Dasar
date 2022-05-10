package main

import "fmt"

func main() {
	var (
		a = "Nurcholis"
		b = "a"
	)

	var result bool = a == b
	fmt.Println(result)

	var (
		i = 100
		j =10
	)

	fmt.Println(i < j)
	fmt.Println(i > j)
	fmt.Println(i == j)
	fmt.Println(i != j)
}