package main

import "fmt"

func main() {
	var name string

	name = "Raisa Supriatna"
	fmt.Println(name)

	name = "Supriatna Raisa"
	fmt.Println(name)

	var phi = 3.14
	fmt.Println("Nilai Phi = ", phi)

	var age int8 = 22
	fmt.Println(age)

	gender := "Male"
	fmt.Println(gender)

	var (
		firstName = "Raisa"
		lastName  = "Supriatna"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
