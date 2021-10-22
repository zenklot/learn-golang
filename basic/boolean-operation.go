package main

import "fmt"

func main() {
	nilaiAkhir := 89
	absen := 77

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsen := absen > 80

	lulus := lulusNilaiAkhir && lulusAbsen

	fmt.Println(lulus)

}
