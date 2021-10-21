package main

import "fmt"

func main() {
	// single constant
	const firstName = "Raisa"
	const lastName = "Supriatna"

	fmt.Println(firstName, lastName)

	// Multi constant
	const (
		age      = 22
		birthday = "12 Nov 1998"
	)

	fmt.Println(age, "thn, ", birthday)
}
