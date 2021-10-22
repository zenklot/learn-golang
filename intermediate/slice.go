package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
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

	var slice1 = months[4:7]
	fmt.Println(slice1)
	// Panjang pakai len
	fmt.Println(len(slice1))
	// Kapasitas
	fmt.Println(cap(slice1))

	// Mengubah array asli : berpengaruh ke aray slice
	months[5] = "Gozenx"
	fmt.Println(slice1)

	// Mengubah array slice : berpengaruh ke array asli
	slice1[1] = "Juni Tai"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	// Melakukan append sehinga menghasilkan data array baru
	var slice3 = append(slice2, "Gozenx")
	fmt.Println(slice3)

	slice3[1] = "Bukan desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// Membuat array baru dengan fungsi : make
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Raisa"
	newSlice[1] = "Supriatna"

	fmt.Println(newSlice)

	// Menyalin array dengan fungsi : copy
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	// ini array
	fmt.Println(iniArray)

	// ini slice
	fmt.Println(iniSlice)

}
