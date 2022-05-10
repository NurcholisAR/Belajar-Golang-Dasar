package main

import "fmt"

func main() {
	var (
		ujian = 80
		absen = 80

		lulusUjian bool  = ujian >= 80
		lulusAbsensi bool = absen >= 80

		lulus = lulusUjian && lulusAbsensi
	)

	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absen >= 80)

}