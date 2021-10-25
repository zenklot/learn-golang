package other

import "fmt"

// Huruf kecil tidak bisa diakses di luar, Hurup besar awalnya bisa diakses
func sayHello(name string) {
	fmt.Println("Hello", name)
}
