package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Raisa",
		"address": "Ciamis",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println("address")

	person["title"] = "Programmer"

	fmt.Println(person["title"])

	// fungi len untuk mendapatkan jumlah data di map
	fmt.Println(len(person))
	// fungsi map[key] mengambil data di map dengan key
	// fungsi map[key] = value mengubah data di map dengan key
	// fungsi make(map[TypeKey]TypeValue) Membuat map baru
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Raisa Supriatna"
	book["upps"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))
	// fungsi delete(map, key) menghapus data di map dengan key
	delete(book, "upps")
	fmt.Println(book)
	fmt.Println(len(book))
}
