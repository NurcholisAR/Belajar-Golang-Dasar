package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApp() {
    //* defer function adalah function yg bisa kita jadwalkan untuk dieksekusi setelah sebuah function dieksekusi
    //* defer function akan selalu dieksekusi walaupun terjadi error di function yg dieksekusi
    //* kita memberitahu go-lang bahwa ketika func runApp() selesai dikesekusi maka harus diakhiri dengan func logging()
    //! biasakan gunakan defer diawal
    defer logging()
    fmt.Println("Aplikasi berjalan")
}


func main() {
    runApp()
}