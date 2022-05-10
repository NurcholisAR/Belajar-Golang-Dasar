package main

import "fmt"

type HasName interface {
	//? returnnya adalah string
	getName() string
	getAge() int
}

func SayHello(hashName HasName) {
	fmt.Println("Hello", hashName.getName())
	fmt.Println(hashName.getAge())
}

//* ketika ingin memanggil function sayhello maka harus mengikuti kontrak si hashname
//? implementasi interface hasname
//? cth 1
type Orang struct {
	Name string
	Age  int
}

func (orang Orang) getName() string {
	return orang.Name
}
func (age Orang) getAge() int {
	return age.Age
}

//? cth 2
type Heman struct {
	Name string
	Age  int
}

func (hewan Heman) getName() string {
	return hewan.Name
}
func (age Heman) getAge() int {
	return age.Age
}

func main() {
	joko := Orang{
		Name: "Joko",
		Age:  25,
	}
	SayHello(joko)

	cat := Heman{
		Name: "Kucing",
		Age:  12,
	}
	SayHello(cat)

}
