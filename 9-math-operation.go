package main

import "fmt"

func main() {
	var (
		a = 10
		b = 10
		c = a + b
		d = 10 * 10
	)

	fmt.Println(c)
	fmt.Println(d)
	
	//* Augmented Assignments
	
	var i = 10
	i+= 10 // i = i + 10

	fmt.Println(i)

	//* Unary Operator

	var j = 10
	j++ // j = j + 1
	var negative = -10
	var positive = +10

	fmt.Println(j)
	fmt.Println(negative)
	fmt.Println(positive)



}