package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

//? tag
type SampleTag struct {
	Name string `required:"true" max:"10"`
}

func IsValid(object interface{}) bool {
	t := reflect.TypeOf(object)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(object).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

type ContohLagi struct {
	Name string `required:"true"`
	Age  int    `required:"true"`
}

func main() {
	sample := Sample{"Budi"}

	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)

	sampleValue := reflect.ValueOf(sample)
	fmt.Println(sampleValue.Field(0))

	//? Tag
	sampleTag := SampleTag{"Contoh"}

	sampleTagType := reflect.TypeOf(sampleTag)
	fmt.Println(sampleTagType.Field(0).Tag.Get("required"))
	fmt.Println(sampleTagType.Field(0).Tag.Get("max"))

	//! jika tidak ada tag required maka hasilnya true, jadi jangan lupa tulis required
	sampleTag.Name = ""
	fmt.Println(IsValid(sampleTag))

	contoh := ContohLagi{"Agus", 20}
	fmt.Println(IsValid(contoh))
}
