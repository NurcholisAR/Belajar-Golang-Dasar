package main

import "fmt"

func main() {
	counter := 0
	val := func() {

		counter := 1
		counter++
		fmt.Println(counter)
	}

	val()
	val()
	fmt.Println(counter)

}
