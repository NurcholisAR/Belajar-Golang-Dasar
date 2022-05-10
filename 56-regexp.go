package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("ePo"))

	fmt.Println(regex.FindAllString("eko,edo,osk", -1))
}
