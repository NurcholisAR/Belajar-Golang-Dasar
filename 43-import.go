package main

import (
	"fmt"
	//* "golang-dasar/createpackage"
	helper "golang-dasar/createpackage"
)

func main() {
	//? cratepackage diambil nama package,atau bisa dikasih nama
	//* result := createpackage.SayHello("budi")
	result := helper.SayHello("eko")
	fmt.Println(result)
}
