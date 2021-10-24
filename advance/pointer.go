package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// 	var address1 Address = Address{"Ciamis", "Jawa Barat", "Indonesia"}
	// 	address2 := &address1

	// 	address2.City = "Bandung"

	// 	fmt.Println(address1)
	// 	fmt.Println(address2)

	var address1 Address = Address{"Ciamis", "Jawa Barat", "Indonesia"}
	address2 := &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	fmt.Println(address4)

}
