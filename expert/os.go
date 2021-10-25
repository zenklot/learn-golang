package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("argument :", args)
	// fmt.Println("argument :", args[1])
	// fmt.Println("argument :", args[2])
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hosname :", name)
	} else {
		fmt.Println("Error ", err.Error())
	}
}
