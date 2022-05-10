package main

import "fmt"

func main() {
	//* tipe data slice adalah potongan dari array
	//* slice mirip dengan array yg membedakan adalah ukuran slice bisa berubah
	//! jika data array diubah maka data slice juga ikut berubah, karena slice reference ke array

	//! Perbedaan Slice dan Array
	iniArray := [3]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	//* jika tidak tau ukuran array maka ditulis ...
	var months = [...]string{
		"January",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	//* membuat slice dari array
	//? array[low:high] -> membuat slice dari array dimulai index low sampai index sebelum high
	var slice = months[4:7]
	fmt.Println(slice)

	//? array[low:] -> membuat slice dari array dimulai index low sampai index terakhir di array
	var slice1 = months[4:]
	fmt.Println(slice1)

	//? array[:high] -> membuat slice dari array dimulai index 0 sampai index sebelum high
	var slice2 = months[:7]
	fmt.Println(slice2)

	//? array[:] -> membuat slice dari array dimulai index 0 sampai index terakhir di array
	var slice3 = months[:]
	fmt.Println(slice3)

	//* function slice
	//? len(slice) -> untuk mendapatkan panjang slice
	fmt.Println(len(slice))
	fmt.Println(len(slice1))
	fmt.Println(len(slice2))
	fmt.Println(len(slice3))

	//? cap(slice) -> untuk mendapatkan kapasitas slice
	fmt.Println(cap(slice))
	fmt.Println(cap(slice1))
	fmt.Println(cap(slice2))
	fmt.Println(cap(slice3))

	//? append(slice, data) -> membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas (array) sudah penuh, maka akan membuat array baru
	//? array masih bisa menampung
	var tambahSlice = append(slice2, "Bulan Baru")

	//? kapasitas array sudah penuh
	var tambahSlice1 = append(slice3, "Bukan Bulan baru")

	fmt.Println(tambahSlice)
	fmt.Println(tambahSlice1)
	fmt.Println(months)

	//? make([]tipe data, length, capacity) -> membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "A"
	newSlice[1] = "B"

	fmt.Println(newSlice)

	//? copy(destination, source) -> menyalin slice dari source ke destination

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
