package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbay := getGoodBye
	fmt.Println(goodbay("gozenx"))
}
