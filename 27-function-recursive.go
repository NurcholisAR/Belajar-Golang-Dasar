package main

import "fmt"

func factorialLoop(value int) int {
	total := 1
	for i := value; i > 0; i-- {
		total *= i
	}
	return total
}

func factorialRecusive(value int) int {
    //* dicek untuk menghentikan perulangan
    switch value {
    case 1:
        return 1
    default:
        return value * factorialRecusive(value -1)
    }
}

func main() {
	fmt.Println(factorialLoop(5))
    fmt.Println(5*5*5*5*5)
    fmt.Println(factorialRecusive(5))
}