package main

import "fmt"

//* recover adalah function yg bisa kita gunakan untuk menangkap data panic,
//* dengan recover proses panic akan berhenti, sehingga program akan tetap berjalan

func endApp1() {
    message := recover()
    if message != nil {
        fmt.Println("Error dengan Message:", message)
    }
	fmt.Println("Aplikasi selesai")
}

func runApp2(error bool) {
    defer endApp1()
    if error {
        panic("Aplikasi Error!")
    }
    //! jangan menulis recover dibawah panic, karena tidak akan dieksekusi.
    //* jika ingin menggunakan recover maka sebaiknya ditulis di defer function
    fmt.Println("Aplikasi berjalan")
}

func main() {
    runApp2(true)
    fmt.Println("Nurcholis")
}