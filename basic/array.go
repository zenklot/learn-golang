package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Raisa"
	names[1] = "Supriatna"
	names[2] = "Raisa Supriatna"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)

	// Fungsi untuk mendapatkan panjang array walaupun tidak ada datanya
	fmt.Println(len(values))
}
