package main

import "fmt"

/*
mirip kayak object di PHP
atau json di JS
tapi ini cuma bikin template data doang, gak bisa langsung ada datanya
jadi datanya perlu di isi setelah di deklarasi struct nya
*/
type Customer struct {
	Name, Address string // kalo typenya sama bisa di kasih typedata nya satu aja di belakang, gak perlu satu satu
	Age           int
}

// struct method
func (customer Customer) sayHello(name string) {
	/*
		jadi function sayHello ini udah berubah jadi method karena function awalnya ngambil dari template struct Customer
		jadi dari function ini yang perlu di lakukan adalah panggil data Customer dulu baru sayHello
		contoh: angga.sayHello(nameString)
	*/
	fmt.Println("Hello,", name, "my name is", customer.Name)
}

func main() {
	var angga Customer
	fmt.Println(angga)

	angga.Name = "Angga"
	angga.Address = "Indonesia"
	angga.Age = 28
	fmt.Println(angga)
	fmt.Println(angga.Name)
	fmt.Println(angga.Address)
	fmt.Println(angga.Age)

	// struct literal
	jokan := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     28,
	}
	fmt.Println(jokan)

	budi := Customer{"Budi", "Berlin", 18}
	fmt.Println(budi)

	angga.sayHello("Kucing")
}
