package main

import "fmt"

func sumAll(numbers ...int) int {

	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {

	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println(total)

	slice := []int{1, 2, 3, 4, 5}
	totalSlice := sumAll(slice...)
	fmt.Println(totalSlice)

}
