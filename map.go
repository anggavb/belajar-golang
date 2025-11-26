package main

import "fmt"

func main() {
	// begini bisa kalo mau di isi manual,
	// karena sifatnya tidak seperti array yang harus di deklarikan kapasitasnya

	//var person = map[string]string{}
	//person["name"] = "Angga"
	//person["address"] = "Bogor"
	//person["age"] = "28"

	person := map[string]string{
		"name":    "Angga",
		"address": "Bogor",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	fmt.Println(person["salah"]) // kalo gaada dia akan tampilkan default value yaitu value kosong ("")

	book := make(map[string]string) // cara lain bikin map, berlaku di slice juga bisa
	book["title"] = "Belajar Golang"
	book["author"] = "Angga"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)

}
