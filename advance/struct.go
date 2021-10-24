package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var raisa Customer
	raisa.Name = "Raisa Supriatna"
	raisa.Address = "Indonesia"
	raisa.Age = 22

	supriatna := Customer{
		Name:    "Supriatna Raisa",
		Address: "Indonesia",
		Age:     22,
	}
	fmt.Println(raisa)
	fmt.Println(supriatna)

}
