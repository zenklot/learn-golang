package main

import "fmt"

func main() {

	// cara satu
	// ==========================
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// Cara kedua
	// ==========================
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"Raisa", "Supriatna", "Gozenx", "Zenklot"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	animals := []string{"dog", "cat", "rabbit", "monkey"}
	for _, value := range animals {
		fmt.Println(value)
	}

	foods := make(map[string]string)
	foods["goreng"] = "Telur"
	foods["kuah"] = "Sayur"
	foods["panggang"] = "Daging"

	for key, val := range foods {
		fmt.Println(key, "=", val)
	}

}
