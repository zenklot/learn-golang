package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.UTC())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)
	layout := time.RFC3339
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
}
