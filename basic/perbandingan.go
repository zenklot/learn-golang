package main

import "fmt"

func main() {
	var name1 = "raisa"
	var name2 = "Raisa"

	result := name1 == name2

	fmt.Println(result)
	fmt.Println(name2 > name1)

	var age1 = 34
	var age2 = 21

	largerAge := age1 > age2

	fmt.Println(largerAge)

	value1 := 100
	value2 := 78

	minValue1 := value1 < value2

	fmt.Println(minValue1)
}
