package main

import "fmt"

//* type assertions merupakan kemampuan merubah tipe data menjadi tipe data yg kita inginkan
//* fitur ini sering sekali digunakan ketika kita bertemu dengan interface kosong
func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	//* melakukan konversi manual / secara langsung
	// var resultToString string = result.(string)
	// fmt.Println(resultToString)

	//! perlu diingat jika melakukan type assertions secara langsung maka pastikan yg dirubah tipe datanya sama, jika salah maka terjadi panic

	//? panic
	// var resultToInt int = result.(int)
	//* output = panic: interface conversion: interface {} is string, not int

	//* untuk lebih aman, gunakan switch expression
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}

}
