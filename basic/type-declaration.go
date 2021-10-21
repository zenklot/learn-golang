package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var noktp NoKTP = "3276567891004"
	fmt.Println(noktp)

	var marriedStatus Married = false
	fmt.Println(marriedStatus)

}
