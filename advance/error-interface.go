package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan NOL")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}
func main() {
	result, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("error", err.Error())
	}
}
