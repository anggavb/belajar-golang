package main

import "fmt"

func main() {
	var name string
	name = "Angga"
	fmt.Println(name)

	name = "Anggaaaa"
	fmt.Println(name)

	var fullName = "Angga VB"
	fmt.Println(fullName)

	simpleName := "Angga VB using simplfy variable"
	fmt.Println(simpleName)

	simpleName = "Angga only"
	fmt.Println(simpleName)

	var (
		firstName = "Angga"
		lastName  = "Vb"
	)
	fmt.Println(firstName, lastName)
}
