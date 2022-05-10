package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	// fmt.Println(time.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, err := time.Parse(layout, "2022-01-02")

	switch err {
	case nil:
		fmt.Println(parse)
	default:
		fmt.Println(err.Error())
	}
}
