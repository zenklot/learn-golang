package main

import "fmt"

type nameFilter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
func sayHelloWithFilter(name string, filter nameFilter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
}
