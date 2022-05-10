package main

import "fmt"

//* struct bisa digunakan sebagai parameter didalam function
type Customer struct {
	Name, Address string
}

//* ditulis sebelum nama function
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}
func (cus Customer) sayHalo() {
	fmt.Println("Hello", cus.Name)
}

func main() {
	var budi Customer
	budi.Name = "Budi"
	//* memanggil function hi
	budi.sayHi("Joko")

	joko := Customer{Name: "Joko"}
	joko.sayHalo()

}
