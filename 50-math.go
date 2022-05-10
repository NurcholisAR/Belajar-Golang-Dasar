package main

import (
	"fmt"
	"math"
)

func main() {
	//* math.Round(float64)  --> membulatkan float64 keatas atau kebawah, sesuai yang paling dekat
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	//* math.Floor(float64) --> membulatkan float64 kebawah
	fmt.Println(math.Floor(1.7))

	//* math.Ceil(float64) --> membulatkan float64 keatas
	fmt.Println(math.Ceil(1.3))

	//* math.Max(float64, float64) --> mengmballikan nilai float64 paling besa
	fmt.Println(math.Max(10, 20))

	//* math.Min(float64, float64) --> mengembalikan nilai float64 paling kecil
	fmt.Println(math.Min(10, 20))
}
