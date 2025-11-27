package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	fmt.Println(address1)

	address2 := &address1

	address2.City = "New York"
	fmt.Println(address1)
	fmt.Println(address2)

	// fungsi *address2 ini untuk mengupdate isi struct sesuai dengan isi nya
	// jadi kalo ada variabel lain yang pake struct ini bakalan ikut juga datanya dengan yang terbaru ini
	*address2 = Address{"Kuningan", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
