package main

import "fmt"

func main() {
	counter := 0
	name := "raisa"

	increment := func() {
		// name = "supriatna"
		name := "supriatna"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
