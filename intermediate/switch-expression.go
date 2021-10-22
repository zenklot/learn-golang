package main

import "fmt"

func main() {
	name := "raisasupriatna"

	switch name {
	case "raisa":
		fmt.Println("Hello Raisa")
	case "supriaatna":
		fmt.Println("Hallo Supriatna")
	default:
		fmt.Println("Haloo semuanya!")
	}

	// Short Statement

	switch length := len(name); length >= 5 {
	case true:
		fmt.Println("Ya nama lebih dari sama dengan 5")
	case false:
		fmt.Println("Nama tidak lebih dari 5")
	}

	length := len(name)
	// Switch tanpa kondisi
	switch {
	case length > 10:
		fmt.Println("Name to long")
	case length > 5:
		fmt.Println("Name good")
	default:
		fmt.Println("Name right")
	}
}
