package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`g([a-z])z`)

	fmt.Println(regex.MatchString("goz"))
	fmt.Println(regex.MatchString("gez"))
	fmt.Println(regex.MatchString("goZ"))

	search := regex.FindAllString("goz zenx gez gaz guz", -1)
	fmt.Println(search)
}
