package createpackage

//* dibahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk access modifier terhadap suatu function atau variable cth: private, public

//*  di golang, untuk menentukan access modifier, cukup dengan nama function atau variable

//*  jika namanya, diawali dengan huruf besar, maka artinya bisa diakses dari package lain, jika diawali dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain.

var verion string = "1.0.0"               //* -> tidak bisa diakses dari luar
var Application string = "Belajar golang" //* -> bisa diakses

//* tidak bisa diakses dari luar
func sayGoodBye(name string) string {
	return "Hello " + name
}

//* bisa diakses
func SayHello1(name string) string {
	return "Hello " + name
}
