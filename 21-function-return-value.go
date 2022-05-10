package main

import "fmt"

func getHello(name string) string {

	switch name {
	case "":
		return "Hello Bro"
	default:
		return "Hello " + name
	}
}

func main() {

	result := getHello("coy")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
