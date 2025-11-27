package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//alamat1 := &Address{}

	/*
		codenya mirip kayak di atas,
		fungsinya adalah deklarasiin variabel ini seperti data langsungnya struct
		jadi kalo ada variabel baru yang pake isi datanya dari variabel new() ini
		maka walaupun si variabel baru ini update maka variabel lama juga akan ikut terupdate
		sebenarnya persis sama seperti pointer &variable
	*/
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.City = "Bogor"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
