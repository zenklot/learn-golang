package main

import "fmt"

func main() {
	a := 10
	b := 66
	c := 5

	// penjumlahan
	fmt.Println(a + b)
	resultAdd := a + c
	fmt.Println(resultAdd)

	// pengurangan
	fmt.Println(b - a)

	// perkalian
	fmt.Println(a * c)

	// pembagian
	fmt.Println(a / c)

	// sisa pembagian
	fmt.Println(b % c)

	// Augmented assignments
	a += 20
	fmt.Println(a)

	// Unary Operator
	// ++ : a=a+1
	// -- : a=a-1
	// -  : Negative
	// +  : Positive
	// !  : Boolean Kebalikan

	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

}
