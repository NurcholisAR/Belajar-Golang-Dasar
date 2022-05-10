package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello,", filter(name))
}

func spamFilter(name string) string {
	switch name {
	case "Anjing":
		return "***"
	default:
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Nurcholis", filter)
}
