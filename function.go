package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHello()
	sayHelloTo("Angga", "Vb")
}
