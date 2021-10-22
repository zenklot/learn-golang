package main

import "fmt"

func main() {

	// Penggunaan Break
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke - ", i)
		if i == 5 {
			break
		}
	}
	fmt.Println("")
	fmt.Println("")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke - ", i)
	}
}
