package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddressToID(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Ciamis", "Jawa Barat", ""}
	changeAddressToID(&address1)

	fmt.Println(address1)
}
