package main

import "fmt"

func main() {
	name := "Raisa"

	if name == "raisa" {
		fmt.Println("Hallo Raisa")
	}

	name2 := "Supriatna"

	if name2 == "raisa" {
		fmt.Println("Hallo Raisa")
	} else {
		fmt.Println("Hi.., Masuk Else")
	}

	if name == "raisa" {
		fmt.Println("Benar Raisa")
	} else if name2 == "supriatna" {
		fmt.Println("Benar ini supriatna")
	} else {
		fmt.Println("Semua Salah")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama sudah ok")
	} else {
		fmt.Println("nama terlalu panjang")
	}
}
