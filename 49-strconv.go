package main

import (
	"fmt"
	"strconv"
)

func main() {
	//* strconv.ParseTipeDataApa()  -> untuk merubah string menjadi tipe data yg kita inginkan

	//* hampir semua strconv bisa mengembalikan error, jadi jangan lupa buat 2 variable

	boolean, err := strconv.ParseBool("false")

	switch err {
	case nil:
		fmt.Println(boolean)
	default:
		fmt.Println(err.Error())
	}

	//? base
	//* bin = 2
	//* oct = 8
	//* dec = 10
	//* hex = 16
	number, err := strconv.ParseInt("1000000", 10, 64)
	switch err {
	case nil:
		fmt.Println(number)
	default:
		fmt.Println(err.Error())
	}

	//? untuk cara gampangnya gunakan Atoi dibanding ParseInt
	valueInt, _ := strconv.Atoi("1000000")
	fmt.Println(valueInt)

	valueString := strconv.Itoa(1000000)
	fmt.Println(valueString)

	//* strconv.FormatTipeDataApa() -> untuk merubah tipe data yg lain menjadi string

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)
}
