package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {

	switch blacklist(name) {
	case true:
		fmt.Println("You are blocked", name)
	case false:
		fmt.Println("Welcome", name)
	}
}

//* jika membuat function manual
// func admin(name string) bool {
// 	return name == "admin"
// }
// func root(name string) bool {
// 	return name == "root"
// }

func main() {
	admin := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", admin)
	registerUser("Nurcholis", admin)

	//* atau bisa juga
	registerUser("root", func(s string) bool {
		return s == "root"
	})
	registerUser("Nurcholis", func(s string) bool {
		return s == "root"
	})
}
