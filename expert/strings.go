package main

import (
	"fmt"
	"strings"
)

func main() {
	hasil := strings.Trim(".....Hilangkan titk di dalam tulisan ini.....", ".")
	fmt.Println(hasil)
	fmt.Println(strings.Contains("Raisa Supriatna", "sup"))
	fmt.Println(strings.Split("Gozenx Supriatna", " "))
	fmt.Println(strings.ToLower("RAISA SupRIatna"))
	fmt.Println(strings.ToUpper("raisa supriatna"))
	fmt.Println(strings.ToTitle("raisa Supriatna"))
	fmt.Println(strings.ReplaceAll("goz goz goz goz", "goz", "zenx"))

}
