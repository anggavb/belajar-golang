package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function return value
func getHello(name string) string {
	return "Hello " + name
}

// returning function values
func getFullName() (string, string) { // kurung pertama adalah parameter, kurung kedua adalah type/jenis variabel dari return
	return "Angga", "Vb"
}

func getFullName2() (firstN, middleN, lastN string) { // kalo sama semua bisa dibuat jadi satu type aja kayak begini
	firstN = "Angga"
	middleN = "-"
	lastN = "Vb"

	return firstN, middleN, lastN
}

func main() {
	sayHello()
	sayHelloTo("Angga", "Vb")

	result := getHello("Angga")
	fmt.Println(result)
	fmt.Println(getHello("Joko"))

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	_, last := getFullName() // kalo gabutuh perlu diganti jadi (_)
	fmt.Println(last)

	a, b, c := getFullName2()
	fmt.Println(a, b, c)
}
