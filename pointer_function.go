package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountry(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{City: "Bandung", Province: "Jawa Barat", Country: "Malaysia"}
	fmt.Println(address)

	/*
		skemanya sama kayak bikin variabel pointer
		var address *Address = &Address{}

		kalo contoh ini adalah data awalnya Country = Malaysia
		nah ketika dikirim ke ChageCountry() dan type struct nya udah ditambahin titik pointer (*)
		maka data yang udah masuk ke struct Address juga akan terupdate
	*/
	ChangeCountry(&address)
	fmt.Println(address)
}
