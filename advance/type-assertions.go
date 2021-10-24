package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("default", value)
	}

	fmt.Println(resultString)
}
