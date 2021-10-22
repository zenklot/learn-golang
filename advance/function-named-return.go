package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Raisa"
	middleName = "Gozenx"
	lastName = "Supriatna"

	return
}

func main() {
	first, middle, last := getCompleteName()
	fmt.Println(first, middle, last)
}
