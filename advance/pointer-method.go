package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	goz := Man{"Gozenx"}
	goz.Married()

	fmt.Println(goz.Name)

}
