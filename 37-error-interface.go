package main

import (
	"errors"
	"fmt"
)

func Pembagian(value int, pembagi int) (int, error) {
    switch pembagi {
    case 0:
        return 0, errors.New("Tidak dapat dibagi 0")
    default:
        result := value / pembagi
        //? karena error adalah interface maka bisa menggunakan nil
        return result, nil

    }
} 

func main() {
    //* contoh pemmbuatan error
	var contohError error = errors.New("Ups error")
    fmt.Println(contohError.Error())

    //? karena return valuenya 2, maka dibuat 2 variable
    result,err := Pembagian(10,0)
    if err == nil {
        fmt.Println("Hasil", result)
    } else {
        fmt.Println("Error :",err.Error())
    }

}