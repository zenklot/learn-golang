package main

import "fmt"

func getFullName() (string, string) {
	return "raisa", "supriatna"
}
func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// Mengignore beberapa value
	nickname, _ := getFullName()
	fmt.Println(nickname)
}
