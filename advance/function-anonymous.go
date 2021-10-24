package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked")
	} else {
		fmt.Println("Welcome")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing", blacklist)

	registerUser("gozenx", func(name string) bool {
		return name == "admin"
	})
}
