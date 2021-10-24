package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
	message := recover()
	if message != nil {
		fmt.Println("Error", message)
	}
}

func runApp(value int) {
	defer logging()
	fmt.Println("Run App")
	result := 10 / value
	fmt.Println(result)
}

func endApp(err bool) {
	defer logging()
	if err {
		panic("gagal menghentikan error")
	}

	fmt.Println("Menghentikan aplikasi")

}

func main() {
	runApp(12)
	endApp(true)
	fmt.Println("Goz")
}
