package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name Is", customer.Name)
}
func main() {
	goz := Customer{
		Name:    "Gozenx",
		Address: "Indonesia",
		Age:     22,
	}
	goz.sayHi("raisa")
	fmt.Println(goz)
}
