package main

import (
	"fmt"
	access "golang-dasar/createpackage"
)

func main() {

	result := access.SayHello1("Joko")
	fmt.Println(access.Application)
	fmt.Println(result)

}
