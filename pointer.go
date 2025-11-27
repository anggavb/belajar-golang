package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "New York" // ini tidak berubah di address1, karena var address2 adalah copy dari address1, mirip PHP
	fmt.Println(address1)
	fmt.Println(address2)

	//var address3 *Address = &address1
	address3 := &address1 // ini namanya pointer, persis kayak di PHP function use &variable
	address3.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
