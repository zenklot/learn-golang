package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hallo bro.."
	} else {
		return "Hello " + name
	}
}
func main() {
	result := getHello("")
	fmt.Println(result)
}
