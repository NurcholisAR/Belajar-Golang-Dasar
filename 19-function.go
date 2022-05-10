package main

import "fmt"

func sayHello() {
	fmt.Println("Hallo")
}

func main() {
	// sayHello()

	for i := 0; i < 5; i++ {
		sayHello()
	}
}
